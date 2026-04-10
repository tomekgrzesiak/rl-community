package rlclient

import (
	"context"
	"fmt"
	"net/url"

	"github.com/rl-community/rl-spectra-assure/internal/domain"
)

// Ensure Client implements the port.
var _ domain.CommunityRepository = (*Client)(nil)

// Search implements domain.CommunityRepository.
func (c *Client) Search(ctx context.Context, items []domain.SearchRequest, params domain.SearchParams) (*domain.SearchResponse, error) {
	path := "/find/packages"
	q := url.Values{}
	if params.Offset > 0 {
		q.Set("offset", fmt.Sprintf("%d", params.Offset))
	}
	if params.Limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", params.Limit))
	}
	if params.Compact {
		q.Set("compact", "true")
	}
	if len(q) > 0 {
		path += "?" + q.Encode()
	}

	var resp domain.SearchResponse
	if err := c.do(ctx, "POST", path, items, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVersionReport implements domain.CommunityRepository.
func (c *Client) GetVersionReport(ctx context.Context, params domain.PackageParams) (*domain.VersionReportResponse, error) {
	path := buildVersionPath(params)
	q := buildArtifactQuery(params)
	if len(q) > 0 {
		path += "?" + q.Encode()
	}

	var resp domain.VersionReportResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPackageDetails implements domain.CommunityRepository.
func (c *Client) GetPackageDetails(ctx context.Context, params domain.PackageParams) (*domain.PackageDetailsResponse, error) {
	path := buildPackagePath(params)
	q := buildArtifactQuery(params)
	if params.MatchPattern != "" {
		q.Set("match_pattern", params.MatchPattern)
	}
	if params.MatchExpression != "" {
		q.Set("match_expression", params.MatchExpression)
	}
	if params.Offset > 0 {
		q.Set("offset", fmt.Sprintf("%d", params.Offset))
	}
	if params.Limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", params.Limit))
	}
	if len(q) > 0 {
		path += "?" + q.Encode()
	}

	var resp domain.PackageDetailsResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// buildVersionPath constructs the URL path for the version report endpoint.
// Template: /report/version/pkg:{community}/{namespace}/{package}@{version}
func buildVersionPath(p domain.PackageParams) string {
	base := fmt.Sprintf("/report/version/pkg:%s/%s/%s", p.Community, p.Namespace, p.Package)
	if p.Version != "" {
		base += "@" + p.Version
	}
	return base
}

// buildPackagePath constructs the URL path for the package details endpoint.
// Template: /report/package/pkg:{community}/{namespace}/{package}@{version}
func buildPackagePath(p domain.PackageParams) string {
	base := fmt.Sprintf("/report/package/pkg:%s/%s/%s", p.Community, p.Namespace, p.Package)
	if p.Version != "" {
		base += "@" + p.Version
	}
	return base
}

// buildArtifactQuery builds query parameters for artifact qualifiers.
func buildArtifactQuery(p domain.PackageParams) url.Values {
	q := url.Values{}
	if p.Artifact != "" {
		q.Set("artifact", p.Artifact)
	}
	if p.ArtifactTag != "" {
		q.Set("artifact_tag", p.ArtifactTag)
	}
	return q
}
