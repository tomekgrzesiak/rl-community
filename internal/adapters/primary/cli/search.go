package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/rl-community/rl-spectra-assure/internal/domain"
)

func newSearchCmd(svcFactory func(*cobra.Command) *services) *cobra.Command {
	var (
		sha1            string
		sha256          string
		matchPattern    string
		matchExpression string
		offset          int
		limit           int
		compact         bool
		outputJSON      bool
	)

	cmd := &cobra.Command{
		Use:   "search [purl]",
		Short: "Search for packages by purl, hash, or version pattern",
		Long: `Search the Spectra Assure Community catalogue.

Examples:
  # Search by purl
  rl-spectra-assure search pkg:pypi/numpy@1.26.0

  # Search by SHA1
  rl-spectra-assure search --sha1 d63932d669fe6da664b4183d8e1d5a33a9492b9f

  # Search by SHA256
  rl-spectra-assure search --sha256 94957715a483aa3b9db60c8f22d5498227d8025c407f982eef81e34f37574ffd

  # Search by version glob pattern
  rl-spectra-assure search pkg:pypi/numpy --match-pattern "1.26.*"

  # Search by version expression
  rl-spectra-assure search pkg:pypi/numpy --match-expression "<= 1.26.0"`,
		Args: func(cmd *cobra.Command, args []string) error {
			hasPurl := len(args) == 1
			hasHash := sha1 != "" || sha256 != ""
			if !hasPurl && !hasHash {
				return fmt.Errorf("purl argument or --sha1/--sha256 flag required")
			}
			if hasPurl && hasHash {
				return fmt.Errorf("purl and hash flags are mutually exclusive")
			}
			if matchPattern != "" && matchExpression != "" {
				return fmt.Errorf("--match-pattern and --match-expression are mutually exclusive")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			svc := svcFactory(cmd)
			ctx := cmd.Context()
			params := domain.SearchParams{Offset: offset, Limit: limit, Compact: compact}

			var resp *domain.SearchResponse
			var err error

			switch {
			case sha1 != "":
				resp, err = svc.search.SearchBySHA1(ctx, sha1, params)
			case sha256 != "":
				resp, err = svc.search.SearchBySHA256(ctx, sha256, params)
			case matchPattern != "":
				resp, err = svc.search.SearchByPattern(ctx, args[0], matchPattern, params)
			case matchExpression != "":
				resp, err = svc.search.SearchByExpression(ctx, args[0], matchExpression, params)
			default:
				resp, err = svc.search.SearchByPURL(ctx, args[0], params)
			}
			if err != nil {
				return err
			}

			return printOutput(cmd, resp, outputJSON)
		},
	}

	cmd.Flags().StringVar(&sha1, "sha1", "", "Search by SHA1 hash")
	cmd.Flags().StringVar(&sha256, "sha256", "", "Search by SHA256 hash")
	cmd.Flags().StringVar(&matchPattern, "match-pattern", "", "Version glob pattern (e.g. 1.26.*)")
	cmd.Flags().StringVar(&matchExpression, "match-expression", "", "Version expression (e.g. \"<= 1.26.0\")")
	cmd.Flags().IntVar(&offset, "offset", 0, "Pagination offset")
	cmd.Flags().IntVar(&limit, "limit", 5, "Maximum number of versions in response")
	cmd.Flags().BoolVar(&compact, "compact", false, "Omit optional fields from response")
	cmd.Flags().BoolVar(&outputJSON, "json", false, "Output raw JSON")

	return cmd
}

// printOutput writes the result either as pretty JSON or a human-friendly summary.
func printOutput(cmd *cobra.Command, v any, asJSON bool) error {
	if asJSON {
		enc := json.NewEncoder(cmd.OutOrStdout())
		enc.SetIndent("", "  ")
		return enc.Encode(v)
	}
	// Default: pretty-print JSON for now; can be extended with table formatters.
	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
