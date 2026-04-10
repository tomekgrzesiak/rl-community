// Package cli provides the cobra command tree for rl-spectra-assure.
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/rl-community/rl-spectra-assure/internal/adapters/secondary/rlclient"
	"github.com/rl-community/rl-spectra-assure/internal/app"
)

// services groups the application-layer services used by all commands.
type services struct {
	search         *app.SearchService
	versionReport  *app.VersionReportService
	packageDetails *app.PackageDetailsService
}

// NewRootCmd builds the root cobra command and wires all sub-commands.
func NewRootCmd() *cobra.Command {
	var (
		token   string
		baseURL string
	)

	root := &cobra.Command{
		Use:   "rl-spectra-assure",
		Short: "CLI for the ReversingLabs Spectra Assure Community API",
		Long: `rl-spectra-assure lets you search for software packages and retrieve
security reports from the ReversingLabs Spectra Assure Community.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if token == "" {
				token = os.Getenv("RL_API_TOKEN")
			}
			if token == "" {
				return fmt.Errorf("API token required: set --token or RL_API_TOKEN env var")
			}
			return nil
		},
	}

	root.PersistentFlags().StringVar(&token, "token", "", "Bearer token for authentication (or set RL_API_TOKEN)")
	root.PersistentFlags().StringVar(&baseURL, "base-url", "", "Override the API base URL (default: Community Free tier)")

	// Build services lazily (after flags are parsed) via a factory closure.
	svcFactory := func(cmd *cobra.Command) *services {
		opts := []rlclient.Option{}
		if baseURL != "" {
			opts = append(opts, rlclient.WithBaseURL(baseURL))
		}
		repo := rlclient.New(token, opts...)
		return &services{
			search:         app.NewSearchService(repo),
			versionReport:  app.NewVersionReportService(repo),
			packageDetails: app.NewPackageDetailsService(repo),
		}
	}

	root.AddCommand(
		newSearchCmd(svcFactory),
		newVersionCmd(svcFactory),
		newPackageCmd(svcFactory),
	)

	return root
}
