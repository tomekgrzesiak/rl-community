package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/rl-community/rl-spectra-assure/internal/domain"
)

func newVersionCmd(svcFactory func(*cobra.Command) *services) *cobra.Command {
	var (
		artifact    string
		artifactTag string
		outputJSON  bool
	)

	cmd := &cobra.Command{
		Use:   "version <community> <package> [version]",
		Short: "Show analysis report for a package version",
		Long: `Retrieve the Spectra Assure analysis report for a specific package version.

Arguments:
  community   Repository type: npm, pypi, gem, nuget, vsx, psgallery
  package     Package name (use namespace/package for scoped packages, e.g. @scope/name)
  version     Package version (optional; defaults to latest)

Examples:
  rl-spectra-assure version pypi numpy 1.26.0
  rl-spectra-assure version npm react 18.2.0
  rl-spectra-assure version npm @scope/package 1.0.0
  rl-spectra-assure version pypi numpy`,
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			community, pkg, version, namespace, err := parsePkgArgs(args)
			if err != nil {
				return err
			}

			svc := svcFactory(cmd)
			ctx := cmd.Context()

			params := domain.PackageParams{
				Community:   domain.Community(community),
				Namespace:   namespace,
				Package:     pkg,
				Version:     version,
				Artifact:    artifact,
				ArtifactTag: artifactTag,
			}

			resp, err := svc.versionReport.GetReport(ctx, params)
			if err != nil {
				return err
			}

			return printOutput(cmd, resp, outputJSON)
		},
	}

	cmd.Flags().StringVar(&artifact, "artifact", "", "Artifact file name qualifier")
	cmd.Flags().StringVar(&artifactTag, "artifact-tag", "", "Artifact tag qualifier")
	cmd.Flags().BoolVar(&outputJSON, "json", false, "Output raw JSON")

	return cmd
}

// parsePkgArgs splits CLI args into community, package, version, and namespace.
// Supports scoped packages like @scope/name passed as a single arg.
func parsePkgArgs(args []string) (community, pkg, version, namespace string, err error) {
	community = args[0]
	pkgArg := args[1]

	// Handle scoped namespace: "@scope/name" or "scope/name"
	for i, c := range pkgArg {
		if c == '/' && i > 0 {
			namespace = pkgArg[:i]
			pkg = pkgArg[i+1:]
			break
		}
	}
	if pkg == "" {
		pkg = pkgArg
	}

	if pkg == "" {
		err = fmt.Errorf("package name cannot be empty")
		return
	}

	if len(args) == 3 {
		version = args[2]
	}
	return
}
