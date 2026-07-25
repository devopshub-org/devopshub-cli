package k8s

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestK8sClientInspect(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
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

	client := &K8sClient{Interface: fakeClient}
	ctx := context.Background()

	summary, err := client.Inspect(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if summary.TotalNodes != 1 {
		t.Errorf("expected 1 node, got %d", summary.TotalNodes)
	}
	if summary.ReadyNodes != 1 {
		t.Errorf("expected 1 ready node, got %d", summary.ReadyNodes)
	}
	if summary.TotalPods != 1 {
		t.Errorf("expected 1 pod, got %d", summary.TotalPods)
	}
	if summary.Status != "Healthy" {
		t.Errorf("expected cluster status 'Healthy', got '%s'", summary.Status)
	}
}

func TestK8sClientListPods(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{Name: "pod-a", Namespace: "default"},
			Status:     corev1.PodStatus{Phase: corev1.PodRunning},
		},
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{Name: "pod-b", Namespace: "kube-system"},
			Status:     corev1.PodStatus{Phase: corev1.PodPending},
		},
	)

	client := &K8sClient{Interface: fakeClient}
	ctx := context.Background()

	allPods, err := client.ListPods(ctx, "")
	if err != nil {
		t.Fatalf("unexpected error listing all pods: %v", err)
	}
	if len(allPods) != 2 {
		t.Errorf("expected 2 pods total, got %d", len(allPods))
	}

	nsPods, err := client.ListPods(ctx, "default")
	if err != nil {
		t.Fatalf("unexpected error listing default pods: %v", err)
	}
	if len(nsPods) != 1 {
		t.Errorf("expected 1 pod in default ns, got %d", len(nsPods))
	}
}

func TestK8sClientListNodes(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&corev1.Node{
			ObjectMeta: metav1.ObjectMeta{Name: "node-control-plane", Labels: map[string]string{"node-role.kubernetes.io/control-plane": ""}},
			Status: corev1.NodeStatus{
				Conditions: []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				},
				NodeInfo: corev1.NodeSystemInfo{KubeletVersion: "v1.31.0"},
			},
		},
	)

	client := &K8sClient{Interface: fakeClient}
	ctx := context.Background()

	nodes, err := client.ListNodes(ctx)
	if err != nil {
		t.Fatalf("unexpected error listing nodes: %v", err)
	}
	if len(nodes) != 1 {
		t.Errorf("expected 1 node, got %d", len(nodes))
	}
	if nodes[0].Role != "control-plane" {
		t.Errorf("expected role 'control-plane', got '%s'", nodes[0].Role)
	}
}
