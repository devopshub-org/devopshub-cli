package cluster

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/devopshub-org/devopshub-cli/pkg/k8s"
	"github.com/spf13/cobra"
)

var kubeconfigPath string

// NewClusterCmd creates the parent 'cluster' command and its subcommands.
func NewClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Inspect and manage Kubernetes clusters",
		Long:  "Commands to inspect cluster health, view pods, node capacity, and filter cluster events.",
	}

	cmd.PersistentFlags().StringVar(&kubeconfigPath, "kubeconfig", "", "Path to kubeconfig file (defaults to KUBECONFIG env or ~/.kube/config)")

	cmd.AddCommand(
		newInspectCmd(),
		newEventsCmd(),
		newPodsCmd(),
		newNodesCmd(),
	)

	return cmd
}

func getClient() (k8s.Client, error) {
	return k8s.NewClient(kubeconfigPath)
}

func newInspectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect",
		Short: "Kubernetes cluster health overview",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := getClient()
			if err != nil {
				return err
			}
			summary, err := client.Inspect(context.Background())
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "METRIC\tVALUE")
			fmt.Fprintf(w, "Cluster Status\t%s\n", summary.Status)
			fmt.Fprintf(w, "Server Version\t%s\n", summary.ServerVersion)
			fmt.Fprintf(w, "Nodes (Ready/Total)\t%d / %d\n", summary.ReadyNodes, summary.TotalNodes)
			fmt.Fprintf(w, "Pods (Running/Total)\t%d / %d\n", summary.RunningPods, summary.TotalPods)
			fmt.Fprintf(w, "Warning Events\t%d\n", summary.WarningEvents)
			return w.Flush()
		},
	}
}

func newEventsCmd() *cobra.Command {
	var namespace string
	var eventType string
	var limit int64

	cmd := &cobra.Command{
		Use:   "events",
		Short: "Recent cluster events with filtering",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := getClient()
			if err != nil {
				return err
			}
			events, err := client.ListEvents(context.Background(), namespace, eventType, limit)
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "TYPE\tREASON\tOBJECT\tMESSAGE")
			for _, e := range events {
				msg := e.Message
				if len(msg) > 60 {
					msg = msg[:57] + "..."
				}
				fmt.Fprintf(w, "%s\t%s\t%s/%s\t%s\n", e.Type, e.Reason, e.InvolvedObject.Kind, e.InvolvedObject.Name, msg)
			}
			return w.Flush()
		},
	}

	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Kubernetes namespace")
	cmd.Flags().StringVar(&eventType, "type", "", "Filter by event type (Warning or Normal)")
	cmd.Flags().Int64Var(&limit, "limit", 50, "Maximum number of events to fetch")

	return cmd
}

func newPodsCmd() *cobra.Command {
	var namespace string

	cmd := &cobra.Command{
		Use:   "pods",
		Short: "Pod status with resource usage",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := getClient()
			if err != nil {
				return err
			}
			pods, err := client.ListPods(context.Background(), namespace)
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "NAMESPACE\tNAME\tSTATUS\tRESTARTS\tCPU REQ\tMEM REQ")
			for _, p := range pods {
				fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\t%s\n", p.Namespace, p.Name, p.Status, p.Restarts, p.CPUReq, p.MemReq)
			}
			return w.Flush()
		},
	}

	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Kubernetes namespace (empty for all namespaces)")

	return cmd
}

func newNodesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "nodes",
		Short: "Node capacity and conditions",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := getClient()
			if err != nil {
				return err
			}
			nodes, err := client.ListNodes(context.Background())
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "NAME\tSTATUS\tROLE\tVERSION\tCPU\tMEMORY")
			for _, n := range nodes {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", n.Name, n.Status, n.Role, n.KubeletVersion, n.CPUCapacity, n.MemoryCapacity)
			}
			return w.Flush()
		},
	}
}
