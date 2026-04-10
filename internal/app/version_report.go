package app

import (
	"context"

	"github.com/rl-community/rl-spectra-assure/internal/domain"
)

// VersionReportService handles version report retrieval use cases.
type VersionReportService struct {
	repo domain.CommunityRepository
}

// NewVersionReportService creates a new VersionReportService.
func NewVersionReportService(repo domain.CommunityRepository) *VersionReportService {
	return &VersionReportService{repo: repo}
}

// GetReport retrieves the analysis report for a specific package version.
func (s *VersionReportService) GetReport(ctx context.Context, params domain.PackageParams) (*domain.VersionReportResponse, error) {
	return s.repo.GetVersionReport(ctx, params)
}
