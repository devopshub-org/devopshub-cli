package cloudcmd

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/devopshub-org/devopshub-cli/pkg/cloud"
	"github.com/spf13/cobra"
)

// NewCloudCmd creates the parent 'cloud' command and its subcommands.
func NewCloudCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cloud",
		Short: "Cloud provider authentication and utility commands",
		Long:  "Check, verify, and inspect credentials for AWS, GCP, and Azure cloud providers.",
	}

	cmd.AddCommand(newAuthCmd())

	return cmd
}

func newAuthCmd() *cobra.Command {
	var providerFilter string

	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Cloud provider credential helper (AWS/GCP/Azure)",
		Long:  "Inspects active authentication status and credential locations for cloud providers.",
		RunE: func(cmd *cobra.Command, args []string) error {
			helper := cloud.NewAuthHelper()
			ctx := context.Background()

			var statuses []cloud.ProviderStatus

			switch providerFilter {
			case "aws":
				statuses = append(statuses, helper.CheckAWS(ctx))
			case "gcp":
				statuses = append(statuses, helper.CheckGCP(ctx))
			case "azure":
				statuses = append(statuses, helper.CheckAzure(ctx))
			case "":
				statuses = append(statuses, helper.CheckAWS(ctx), helper.CheckGCP(ctx), helper.CheckAzure(ctx))
			default:
				return fmt.Errorf("unsupported cloud provider '%s' (must be aws, gcp, or azure)", providerFilter)
			}

			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "PROVIDER\tAUTHENTICATED\tIDENTITY\tDETAILS")
			for _, s := range statuses {
				authStr := "No"
				if s.Authenticated {
					authStr = "Yes"
				}
				details := s.Details
				if len(details) > 50 {
					details = details[:47] + "..."
				}
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", s.Provider, authStr, s.Identity, details)
			}
			return w.Flush()
		},
	}

	cmd.Flags().StringVarP(&providerFilter, "provider", "p", "", "Filter by cloud provider (aws, gcp, azure)")

	return cmd
}
