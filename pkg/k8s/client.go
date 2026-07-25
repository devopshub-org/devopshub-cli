package k8s

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Client defines the interface for interacting with Kubernetes clusters.
type Client interface {
	Inspect(ctx context.Context) (*ClusterHealthSummary, error)
	ListEvents(ctx context.Context, namespace, eventType string, limit int64) ([]corev1.Event, error)
	ListPods(ctx context.Context, namespace string) ([]PodSummary, error)
	ListNodes(ctx context.Context) ([]NodeSummary, error)
}

type K8sClient struct {
	Interface kubernetes.Interface
}

type ClusterHealthSummary struct {
	ServerVersion string
	TotalNodes    int
	ReadyNodes    int
	TotalPods     int
	RunningPods   int
	WarningEvents int
	Status        string
}

type PodSummary struct {
	Namespace string
	Name      string
	Status    string
	Restarts  int32
	CPUReq    string
	MemReq    string
	Age       string
}

type NodeSummary struct {
	Name           string
	Status         string
	Role           string
	KubeletVersion string
	CPUCapacity    string
	MemoryCapacity string
	MemoryPressure bool
	DiskPressure   bool
	PIDPressure    bool
}

// NewClient creates a new Kubernetes client using the active kubeconfig.
func NewClient(kubeconfigPath string) (*K8sClient, error) {
	if kubeconfigPath == "" {
		kubeconfigPath = os.Getenv("KUBECONFIG")
	}
	if kubeconfigPath == "" {
		home, err := os.UserHomeDir()
		if err == nil {
			kubeconfigPath = filepath.Join(home, ".kube", "config")
		}
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load kubeconfig from '%s': %w", kubeconfigPath, err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes clientset: %w", err)
	}

	return &K8sClient{Interface: clientset}, nil
}

func (c *K8sClient) Inspect(ctx context.Context) (*ClusterHealthSummary, error) {
	nodes, err := c.Interface.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list nodes: %w", err)
	}

	readyNodes := 0
	for _, n := range nodes.Items {
		for _, cond := range n.Status.Conditions {
			if cond.Type == corev1.NodeReady && cond.Status == corev1.ConditionTrue {
				readyNodes++
			}
		}
	}

	pods, err := c.Interface.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %w", err)
	}

	runningPods := 0
	for _, p := range pods.Items {
		if p.Status.Phase == corev1.PodRunning {
			runningPods++
		}
	}

	events, err := c.Interface.CoreV1().Events("").List(ctx, metav1.ListOptions{})
	warningEvents := 0
	if err == nil {
		for _, e := range events.Items {
			if e.Type == corev1.EventTypeWarning {
				warningEvents++
			}
		}
	}

	status := "Healthy"
	if readyNodes < len(nodes.Items) || warningEvents > 10 {
		status = "Degraded"
	}

	version := "v1.31.0"
	serverVersion, err := c.Interface.Discovery().ServerVersion()
	if err == nil && serverVersion != nil {
		version = serverVersion.GitVersion
	}

	return &ClusterHealthSummary{
		ServerVersion: version,
		TotalNodes:    len(nodes.Items),
		ReadyNodes:    readyNodes,
		TotalPods:     len(pods.Items),
		RunningPods:   runningPods,
		WarningEvents: warningEvents,
		Status:        status,
	}, nil
}

func (c *K8sClient) ListEvents(ctx context.Context, namespace, eventType string, limit int64) ([]corev1.Event, error) {
	opts := metav1.ListOptions{Limit: limit}
	events, err := c.Interface.CoreV1().Events(namespace).List(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list events: %w", err)
	}

	if eventType == "" {
		return events.Items, nil
	}

	filtered := make([]corev1.Event, 0)
	for _, e := range events.Items {
		if e.Type == eventType {
			filtered = append(filtered, e)
		}
	}
	return filtered, nil
}

func (c *K8sClient) ListPods(ctx context.Context, namespace string) ([]PodSummary, error) {
	pods, err := c.Interface.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %w", err)
	}

	summaries := make([]PodSummary, 0, len(pods.Items))
	for _, p := range pods.Items {
		restarts := int32(0)
		for _, cs := range p.Status.ContainerStatuses {
			restarts += cs.RestartCount
		}

		cpuReq := "0m"
		memReq := "0Mi"
		if len(p.Spec.Containers) > 0 {
			if cpu, ok := p.Spec.Containers[0].Resources.Requests[corev1.ResourceCPU]; ok {
				cpuReq = cpu.String()
			}
			if mem, ok := p.Spec.Containers[0].Resources.Requests[corev1.ResourceMemory]; ok {
				memReq = mem.String()
			}
		}

		summaries = append(summaries, PodSummary{
			Namespace: p.Namespace,
			Name:      p.Name,
			Status:    string(p.Status.Phase),
			Restarts:  restarts,
			CPUReq:    cpuReq,
			MemReq:    memReq,
			Age:       p.CreationTimestamp.Format("2006-01-02 15:04"),
		})
	}
	return summaries, nil
}

func (c *K8sClient) ListNodes(ctx context.Context) ([]NodeSummary, error) {
	nodes, err := c.Interface.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list nodes: %w", err)
	}

	summaries := make([]NodeSummary, 0, len(nodes.Items))
	for _, n := range nodes.Items {
		status := "NotReady"
		memPress := false
		diskPress := false
		pidPress := false

		for _, cond := range n.Status.Conditions {
			if cond.Type == corev1.NodeReady && cond.Status == corev1.ConditionTrue {
				status = "Ready"
			}
			if cond.Type == corev1.NodeMemoryPressure && cond.Status == corev1.ConditionTrue {
				memPress = true
			}
			if cond.Type == corev1.NodeDiskPressure && cond.Status == corev1.ConditionTrue {
				diskPress = true
			}
			if cond.Type == corev1.NodePIDPressure && cond.Status == corev1.ConditionTrue {
				pidPress = true
			}
		}

		role := "worker"
		if _, ok := n.Labels["node-role.kubernetes.io/control-plane"]; ok {
			role = "control-plane"
		} else if _, ok := n.Labels["node-role.kubernetes.io/master"]; ok {
			role = "master"
		}

		summaries = append(summaries, NodeSummary{
			Name:           n.Name,
			Status:         status,
			Role:           role,
			KubeletVersion: n.Status.NodeInfo.KubeletVersion,
			CPUCapacity:    n.Status.Capacity.Cpu().String(),
			MemoryCapacity: n.Status.Capacity.Memory().String(),
			MemoryPressure: memPress,
			DiskPressure:   diskPress,
			PIDPressure:    pidPress,
		})
	}
	return summaries, nil
}
