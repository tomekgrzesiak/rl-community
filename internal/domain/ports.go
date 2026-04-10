package domain

import "context"

// CommunityRepository is the driven port for the Spectra Assure Community API.
// Implementations live in adapters/secondary.
type CommunityRepository interface {
	// Search searches for one or more packages by purl, hash, or pattern.
	Search(ctx context.Context, items []SearchRequest, params SearchParams) (*SearchResponse, error)

	// GetVersionReport retrieves the full analysis report for a specific package version.
	GetVersionReport(ctx context.Context, params PackageParams) (*VersionReportResponse, error)

	// GetPackageDetails retrieves metadata and version history for a package.
	GetPackageDetails(ctx context.Context, params PackageParams) (*PackageDetailsResponse, error)
}
