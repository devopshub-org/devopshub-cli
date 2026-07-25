package k8s_test

import (
	"context"
	"testing"

	"github.com/devopshub-org/devopshub-cli/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestK8sClientInspect(t *testing.T) {
	fakeClientset := fake.NewSimpleClientset(
		&corev1.Node{
			ObjectMeta: metav1.ObjectMeta{Name: "node-1"},
			Status: corev1.NodeStatus{
				Conditions: []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				},
			},
		},
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{Name: "pod-1", Namespace: "default"},
			Status:     corev1.PodStatus{Phase: corev1.PodRunning},
		},
	)

	client := &k8s.K8sClient{Interface: fakeClientset}
	summary, err := client.Inspect(context.Background())
	if err != nil {
		t.Fatalf("unexpected error during inspect: %v", err)
	}

	if summary.TotalNodes != 1 {
		t.Errorf("expected 1 total node, got %d", summary.TotalNodes)
	}
	if summary.ReadyNodes != 1 {
		t.Errorf("expected 1 ready node, got %d", summary.ReadyNodes)
	}
	if summary.RunningPods != 1 {
		t.Errorf("expected 1 running pod, got %d", summary.RunningPods)
	}
	if summary.Status != "Healthy" {
		t.Errorf("expected 'Healthy', got '%s'", summary.Status)
	}
}

func TestK8sClientListPods(t *testing.T) {
	fakeClientset := fake.NewSimpleClientset(
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{Name: "nginx", Namespace: "prod"},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name: "nginx",
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("100m"),
								corev1.ResourceMemory: resource.MustParse("128Mi"),
							},
						},
					},
				},
			},
			Status: corev1.PodStatus{
				Phase: corev1.PodRunning,
				ContainerStatuses: []corev1.ContainerStatus{
					{RestartCount: 2},
				},
			},
		},
	)

	client := &k8s.K8sClient{Interface: fakeClientset}
	pods, err := client.ListPods(context.Background(), "prod")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(pods) != 1 {
		t.Fatalf("expected 1 pod, got %d", len(pods))
	}
	if pods[0].Name != "nginx" {
		t.Errorf("expected pod name 'nginx', got '%s'", pods[0].Name)
	}
	if pods[0].Restarts != 2 {
		t.Errorf("expected 2 restarts, got %d", pods[0].Restarts)
	}
}

func TestK8sClientListNodes(t *testing.T) {
	fakeClientset := fake.NewSimpleClientset(
		&corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name:   "master-node",
				Labels: map[string]string{"node-role.kubernetes.io/control-plane": ""},
			},
			Status: corev1.NodeStatus{
				NodeInfo: corev1.NodeSystemInfo{KubeletVersion: "v1.31.0"},
				Conditions: []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				},
				Capacity: corev1.ResourceList{
					corev1.ResourceCPU:    resource.MustParse("4"),
					corev1.ResourceMemory: resource.MustParse("16Gi"),
				},
			},
		},
	)

	client := &k8s.K8sClient{Interface: fakeClientset}
	nodes, err := client.ListNodes(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(nodes))
	}
	if nodes[0].Role != "control-plane" {
		t.Errorf("expected role 'control-plane', got '%s'", nodes[0].Role)
	}
	if nodes[0].Status != "Ready" {
		t.Errorf("expected status 'Ready', got '%s'", nodes[0].Status)
	}
}
