package dockerhub

import (
	"context"
	"fmt"
	"net/http"
)

// RepositoriesService handles communication with the repository
// related methods of the Dockerhub API.
type RepositoriesService service

// RepositoryPatch represents payload to patch a Repository.
type RepositoryPatch struct {
	FullDescription string `json:"full_description,omitempty"`
	Description     string `json:"description,omitempty"`
}

// RepositoryPermissions specifies the permissions of the requesting user
// to the given Repository.
type RepositoryPermissions struct {
	Read  bool `type:"read"`
	Write bool `type:"Write"`
	Admin bool `type:"Admin"`
}

// Repository represents a Dockerhub repository.
type Repository struct {
	User            string  `json:"user"`
	Name            string  `json:"name"`
	Namespace       string  `json:"namespace"`
	RepositoryType  string  `json:"repository_type"`
	Status          int     `json:"status"`
	Description     string  `json:"description"`
	IsPrivate       bool    `json:"is_private"`
	IsAutomated     bool    `json:"is_automated"`
	CanEdit         bool    `json:"can_edit"`
	StarCount       int     `json:"star_count"`
	PullCount       int     `json:"pull_count"`
	LastUpdated     string  `json:"last_updated"`
	IsMigrated      bool    `json:"is_migrated"`
	HasStarred      bool    `json:"has_starred"`
	FullDescription string  `json:"full_description"`
	Affiliation     *string `json:"affiliation"`

	Permissions RepositoryPermissions `json:"repository_permissions"`
}

func (s RepositoriesService) buildRepoSlug(namespace, repo string) string {
	return fmt.Sprintf("/repositories/%s/%s/", namespace, repo)
}

// EditRepository updates a repository.
func (s *RepositoriesService) EditRepository(ctx context.Context, namespace, repo string, patch *RepositoryPatch) (*Repository, error) {
	slug := s.buildRepoSlug(namespace, repo)
	req, err := s.client.NewRequest(http.MethodPatch, slug, patch)
	if err != nil {
		return nil, err
	}

	res := &Repository{}
	if err := s.client.Do(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetRepository gets details for a given repository.
func (s *RepositoriesService) GetRepository(ctx context.Context, namespace, repo string) (*Repository, error) {
	slug := s.buildRepoSlug(namespace, repo)
	req, err := s.client.NewRequest(http.MethodGet, slug, nil)
	if err != nil {
		return nil, err
	}

	res := &Repository{}
	if err := s.client.Do(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
