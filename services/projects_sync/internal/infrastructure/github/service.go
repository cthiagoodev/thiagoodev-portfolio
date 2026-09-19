package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
)

type GithubService interface {
	FetchRepositories(ctx context.Context) ([]entities.Project, error)
}

type GithubServiceImpl struct {
	client  *http.Client
	baseURL *url.URL
}

func NewGithubServiceImpl(client *http.Client, baseURL *url.URL) *GithubServiceImpl {
	return &GithubServiceImpl{
		client:  client,
		baseURL: baseURL,
	}
}

func (s *GithubServiceImpl) FetchRepositories(ctx context.Context) ([]entities.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	const perPage = 100
	projects := make([]Project, 0)

	for page := 1; ; page++ {
		endpoint := s.baseURL.JoinPath("users", "cthiagoodev", "repos")
		query := endpoint.Query()
		query.Set("per_page", strconv.Itoa(perPage))
		query.Set("page", strconv.Itoa(page))
		endpoint.RawQuery = query.Encode()

		pageProjects, err := s.fetchRepositoriesPage(ctx, endpoint)
		if err != nil {
			return nil, fmt.Errorf("fetch GitHub repositories page %d: %w", page, err)
		}

		projects = append(projects, pageProjects...)
		if len(pageProjects) < perPage {
			return MapGithubProjectsToEntities(projects), nil
		}
	}
}

func (s *GithubServiceImpl) fetchRepositoriesPage(ctx context.Context, endpoint *url.URL) ([]Project, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	response, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	var projects []Project
	if err := json.Unmarshal(body, &projects); err != nil {
		return nil, fmt.Errorf("decode response body: %w", err)
	}

	return projects, nil
}
