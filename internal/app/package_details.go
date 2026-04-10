package app

import (
	"context"

	"github.com/rl-community/rl-spectra-assure/internal/domain"
)

// PackageDetailsService handles package metadata retrieval use cases.
type PackageDetailsService struct {
	repo domain.CommunityRepository
}

// NewPackageDetailsService creates a new PackageDetailsService.
func NewPackageDetailsService(repo domain.CommunityRepository) *PackageDetailsService {
	return &PackageDetailsService{repo: repo}
}

// GetDetails retrieves metadata and version history for a package.
func (s *PackageDetailsService) GetDetails(ctx context.Context, params domain.PackageParams) (*domain.PackageDetailsResponse, error) {
	return s.repo.GetPackageDetails(ctx, params)
}
