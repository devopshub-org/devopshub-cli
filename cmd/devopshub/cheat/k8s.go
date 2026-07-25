package cheat

import (
	"github.com/spf13/cobra"

	"github.com/devopshub-org/devopshub-cli/internal/output"
)

// k8sEntry represents a single kubectl cheatsheet entry.
type k8sEntry struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// k8sSection groups related kubectl commands.
type k8sSection struct {
	Title   string     `json:"title"`
	Entries []k8sEntry `json:"entries"`
}

// k8sCheatsheet is the canonical Kubernetes cheatsheet embedded in the binary.
var k8sCheatsheet = []k8sSection{
	{
		Title: "Cluster & Context",
		Entries: []k8sEntry{
			{"kubectl config get-contexts", "List all available contexts"},
			{"kubectl config current-context", "Show current context"},
			{"kubectl config use-context <ctx>", "Switch to a context"},
			{"kubectl cluster-info", "Display cluster endpoint info"},
			{"kubectl version --short", "Show client and server versions"},
			{"kubectl api-resources", "List all available API resources"},
		},
	},
	{
		Title: "Pods",
		Entries: []k8sEntry{
			{"kubectl get pods -A", "List all pods in all namespaces"},
			{"kubectl get pods -n <ns>", "List pods in a specific namespace"},
			{"kubectl get pod <pod> -o yaml", "Get full YAML definition of a pod"},
			{"kubectl describe pod <pod>", "Show detailed pod events and status"},
			{"kubectl logs <pod>", "Print pod logs"},
			{"kubectl logs -f <pod>", "Follow pod logs in real-time"},
			{"kubectl logs <pod> -c <container>", "Logs from a specific container"},
			{"kubectl exec -it <pod> -- /bin/sh", "Open a shell in a running pod"},
			{"kubectl delete pod <pod>", "Delete a pod (will restart if managed)"},
			{"kubectl delete pod <pod> --grace-period=0 --force", "Force-delete a stuck pod"},
		},
	},
	{
		Title: "Deployments",
		Entries: []k8sEntry{
			{"kubectl get deployments -n <ns>", "List deployments in a namespace"},
			{"kubectl describe deployment <name>", "Show deployment details"},
			{"kubectl rollout status deployment/<name>", "Watch rollout progress"},
			{"kubectl rollout history deployment/<name>", "Show rollout history"},
			{"kubectl rollout undo deployment/<name>", "Rollback to previous revision"},
			{"kubectl scale deployment <name> --replicas=3", "Scale deployment replicas"},
			{"kubectl set image deployment/<name> <c>=<image>:<tag>", "Update container image"},
			{"kubectl delete deployment <name>", "Delete a deployment"},
		},
	},
	{
		Title: "Services & Networking",
		Entries: []k8sEntry{
			{"kubectl get svc -A", "List all services"},
			{"kubectl describe svc <name>", "Show service details"},
			{"kubectl port-forward svc/<name> 8080:80", "Forward local port to service"},
			{"kubectl port-forward pod/<pod> 8080:8080", "Forward local port to pod"},
			{"kubectl get ingress -A", "List all ingresses"},
			{"kubectl describe ingress <name>", "Show ingress configuration"},
		},
	},
	{
		Title: "ConfigMaps & Secrets",
		Entries: []k8sEntry{
			{"kubectl get configmaps -n <ns>", "List ConfigMaps"},
			{"kubectl describe configmap <name>", "Show ConfigMap data"},
			{"kubectl get secrets -n <ns>", "List Secrets"},
			{"kubectl get secret <name> -o jsonpath='{.data}'", "Get base64-encoded secret data"},
			{"kubectl create secret generic <n> --from-literal=key=val", "Create a generic secret"},
			{"kubectl create configmap <n> --from-file=./config.yaml", "Create ConfigMap from file"},
		},
	},
	{
		Title: "Nodes",
		Entries: []k8sEntry{
			{"kubectl get nodes", "List cluster nodes"},
			{"kubectl describe node <name>", "Detailed node info"},
			{"kubectl top nodes", "Show node resource usage (requires metrics-server)"},
			{"kubectl cordon <node>", "Mark node as unschedulable"},
			{"kubectl uncordon <node>", "Mark node as schedulable"},
			{"kubectl drain <node> --ignore-daemonsets", "Evict pods before maintenance"},
		},
	},
	{
		Title: "Namespace Management",
		Entries: []k8sEntry{
			{"kubectl get namespaces", "List all namespaces"},
			{"kubectl create namespace <name>", "Create a namespace"},
			{"kubectl delete namespace <name>", "Delete a namespace and all resources"},
			{"kubectl config set-context --current --namespace=<ns>", "Set default namespace"},
		},
	},
	{
		Title: "Apply & Manage Resources",
		Entries: []k8sEntry{
			{"kubectl apply -f <file.yaml>", "Apply a manifest file"},
			{"kubectl apply -f <directory>/", "Apply all manifests in a directory"},
			{"kubectl apply -k <kustomization-dir>/", "Apply via Kustomize"},
			{"kubectl delete -f <file.yaml>", "Delete resources defined in a file"},
			{"kubectl diff -f <file.yaml>", "Show diff before applying"},
			{"kubectl get all -n <ns>", "List all resources in a namespace"},
			{"kubectl get events -n <ns> --sort-by=.lastTimestamp", "View recent events, sorted"},
		},
	},
	{
		Title: "Debugging",
		Entries: []k8sEntry{
			{"kubectl run debug --image=busybox -it --rm -- sh", "Run ephemeral debug pod"},
			{"kubectl debug node/<node> -it --image=ubuntu", "Debug a node directly"},
			{"kubectl get pod <pod> -o jsonpath='{.status.phase}'", "Get pod phase"},
			{"kubectl top pods -n <ns>", "Show pod CPU/memory usage"},
			{"kubectl events --field-selector reason=BackOff", "Filter events by reason"},
		},
	},
}

// k8sSubCmd implements `devopshub cheat k8s`.
type k8sSubCmd struct {
	cmd     *cobra.Command
	printer *output.Printer
}

func newK8sCmd(printer *output.Printer) *k8sSubCmd {
	k := &k8sSubCmd{printer: printer}

	k.cmd = &cobra.Command{
		Use:     "k8s",
		Short:   "Kubernetes (kubectl) cheatsheet",
		Aliases: []string{"kubernetes", "kubectl"},
		Long: `A curated, categorized cheatsheet of essential kubectl commands.

Covers cluster management, pods, deployments, services, configmaps,
secrets, nodes, namespaces, and debugging.

Examples:
  devopshub cheat k8s
  devopshub cheat kubernetes
  devopshub cheat k8s --json`,
		Args: cobra.NoArgs,
		RunE: k.run,
	}

	return k
}

func (k *k8sSubCmd) run(_ *cobra.Command, _ []string) error {
	p := k.printer

	if p.JSONMode() {
		return p.JSON(k8sCheatsheet)
	}

	p.Header("Kubernetes (kubectl) Cheatsheet")
	p.Println("")

	for _, section := range k8sCheatsheet {
		p.Section(section.Title)
		for _, entry := range section.Entries {
			p.Command(entry.Command, entry.Description)
		}
	}

	p.Println("")
	p.Info("Tip: Use --json to get machine-readable output")
	return nil
}
