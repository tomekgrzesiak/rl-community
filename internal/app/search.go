package app

import (
	"context"

	"github.com/rl-community/rl-spectra-assure/internal/domain"
)

// SearchService handles package search use cases.
type SearchService struct {
	repo domain.CommunityRepository
}

// NewSearchService creates a new SearchService.
func NewSearchService(repo domain.CommunityRepository) *SearchService {
	return &SearchService{repo: repo}
}

// SearchByPURL searches for a package by its purl.
func (s *SearchService) SearchByPURL(ctx context.Context, purl string, params domain.SearchParams) (*domain.SearchResponse, error) {
	items := []domain.SearchRequest{
		{UUID: "1", PURL: purl},
	}
	return s.repo.Search(ctx, items, params)
}

// SearchBySHA1 searches for packages containing a SHA1 hash.
func (s *SearchService) SearchBySHA1(ctx context.Context, sha1 string, params domain.SearchParams) (*domain.SearchResponse, error) {
	items := []domain.SearchRequest{
		{UUID: "1", SHA1: sha1},
	}
	return s.repo.Search(ctx, items, params)
}

// SearchBySHA256 searches for packages containing a SHA256 hash.
func (s *SearchService) SearchBySHA256(ctx context.Context, sha256 string, params domain.SearchParams) (*domain.SearchResponse, error) {
	items := []domain.SearchRequest{
		{UUID: "1", SHA256: sha256},
	}
	return s.repo.Search(ctx, items, params)
}

// SearchByPattern searches for package versions matching a glob pattern.
func (s *SearchService) SearchByPattern(ctx context.Context, purl, pattern string, params domain.SearchParams) (*domain.SearchResponse, error) {
	items := []domain.SearchRequest{
		{UUID: "1", PURL: purl, MatchPattern: pattern},
	}
	return s.repo.Search(ctx, items, params)
}

// SearchByExpression searches for package versions matching a version expression.
func (s *SearchService) SearchByExpression(ctx context.Context, purl, expression string, params domain.SearchParams) (*domain.SearchResponse, error) {
	items := []domain.SearchRequest{
		{UUID: "1", PURL: purl, MatchExpression: expression},
	}
	return s.repo.Search(ctx, items, params)
}
