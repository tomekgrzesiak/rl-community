package cli

import (
	"github.com/spf13/cobra"

	"github.com/rl-community/rl-spectra-assure/internal/domain"
)

func newPackageCmd(svcFactory func(*cobra.Command) *services) *cobra.Command {
	var (
		artifact        string
		artifactTag     string
		matchPattern    string
		matchExpression string
		offset          int
		limit           int
		outputJSON      bool
	)

	cmd := &cobra.Command{
		Use:   "package <community> <package> [version]",
		Short: "Show package metadata and version history",
		Long: `Retrieve metadata and version list for a software package.

Arguments:
  community   Repository type: npm, pypi, gem, nuget, vsx, psgallery
  package     Package name (use namespace/package for scoped packages)
  version     Specific version (optional)

Examples:
  rl-spectra-assure package pypi numpy
  rl-spectra-assure package pypi numpy 1.26.0
  rl-spectra-assure package npm react --match-pattern "18.*"
  rl-spectra-assure package npm @scope/package`,
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			community, pkg, version, namespace, err := parsePkgArgs(args)
			if err != nil {
				return err
			}

			svc := svcFactory(cmd)
			ctx := cmd.Context()

			params := domain.PackageParams{
				Community:       domain.Community(community),
				Namespace:       namespace,
				Package:         pkg,
				Version:         version,
				Artifact:        artifact,
				ArtifactTag:     artifactTag,
				MatchPattern:    matchPattern,
				MatchExpression: matchExpression,
				Offset:          offset,
				Limit:           limit,
			}

			resp, err := svc.packageDetails.GetDetails(ctx, params)
			if err != nil {
				return err
			}

			return printOutput(cmd, resp, outputJSON)
		},
	}

	cmd.Flags().StringVar(&artifact, "artifact", "", "Artifact file name qualifier")
	cmd.Flags().StringVar(&artifactTag, "artifact-tag", "", "Artifact tag qualifier")
	cmd.Flags().StringVar(&matchPattern, "match-pattern", "", "Version glob pattern (e.g. 1.26.*)")
	cmd.Flags().StringVar(&matchExpression, "match-expression", "", "Version expression (e.g. \"<= 1.26.0\")")
	cmd.Flags().IntVar(&offset, "offset", 0, "Pagination offset for version list")
	cmd.Flags().IntVar(&limit, "limit", 5, "Maximum number of versions in response")
	cmd.Flags().BoolVar(&outputJSON, "json", false, "Output raw JSON")

	return cmd
}
