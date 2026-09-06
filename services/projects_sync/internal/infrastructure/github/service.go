package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GithubService interface {
	FetchRepositories() ([]Project, error)
}

type GithubServiceImpl struct {
	client  *http.Client
	baseURL *url.URL
}

func NewGithubServiceImpl(client *http.Client, baseURL *url.URL) *GithubServiceImpl {
	return &GithubServiceImpl{
		client,
		baseURL,
	}
}

func (s *GithubServiceImpl) FetchRepositories() ([]Project, error) {
	path := s.baseURL.String() + "/users/cthiagoodev/repos"
	response, err := s.client.Get(path)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	bytes, rErr := io.ReadAll(response.Body)

	if rErr != nil {
		return nil, rErr
	}

	projects := make([]Project, 0)

	jErr := json.Unmarshal(bytes, &projects)

	if jErr != nil {
		return nil, jErr
	}

	return projects, nil
}
