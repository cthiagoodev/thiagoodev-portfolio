package github

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGithubServiceImpl_FetchRepositories(t *testing.T) {
	t.Run("should fetch repositories successfully", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/users/cthiagoodev/repos", r.URL.Path)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				_, _ = w.Write([]byte(`[
					{
						"id": 123,
						"name": "portfolio",
						"html_url": "https://github.com/cthiagoodev/portfolio"
					}
				]`))
			}),
		)

		defer server.Close()

		baseURL, err := url.Parse(server.URL)
		require.NoError(t, err)

		ctx := context.TODO()
		service := NewGithubServiceImpl(
			server.Client(),
			baseURL,
		)

		result, err := service.FetchRepositories(ctx)

		require.NoError(t, err)
		require.Len(t, result, 1)

		assert.Equal(t, int64(123), result[0].Id)
		assert.Equal(t, "portfolio", result[0].Name)
		assert.Equal(
			t,
			"https://github.com/cthiagoodev/portfolio",
			result[0].HtmlUrl,
		)
	})

	t.Run("should return error when github returns non 200", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}),
		)
		defer server.Close()

		baseURL, err := url.Parse(server.URL)
		require.NoError(t, err)

		ctx := context.TODO()
		service := NewGithubServiceImpl(
			server.Client(),
			baseURL,
		)

		result, err := service.FetchRepositories(ctx)

		require.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(
			t,
			err,
			"fetch GitHub repositories page 1: unexpected HTTP status 500",
		)
	})

	t.Run("should return error when response contains invalid json", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)

				_, _ = w.Write([]byte(`invalid-json`))
			}),
		)
		defer server.Close()

		baseURL, err := url.Parse(server.URL)
		require.NoError(t, err)

		ctx := context.TODO()
		service := NewGithubServiceImpl(
			server.Client(),
			baseURL,
		)

		result, err := service.FetchRepositories(ctx)

		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when request fails", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		)

		baseURL, err := url.Parse(server.URL)
		require.NoError(t, err)

		client := server.Client()

		server.Close()

		ctx := context.TODO()
		service := NewGithubServiceImpl(
			client,
			baseURL,
		)

		result, err := service.FetchRepositories(ctx)

		require.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestFetchRepositoriesPagination(t *testing.T) {
	for _, failNextPage := range []bool{false, true} {
		name := "collect all pages"
		if failNextPage {
			name = "discard partial results on failure"
		}
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "100", r.URL.Query().Get("per_page"))
				assert.Equal(t, "/users/cthiagoodev/repos", r.URL.Path)
				switch r.URL.Query().Get("page") {
				case "1":
					projects := make([]Project, 100)
					for i := range projects {
						projects[i].Id = int64(i + 1)
					}
					_ = json.NewEncoder(w).Encode(projects)
				case "2":
					if failNextPage {
						w.WriteHeader(http.StatusBadGateway)
						return
					}
					_ = json.NewEncoder(w).Encode([]Project{{Id: 101}})
				default:
					t.Error("unexpected page")
					w.WriteHeader(http.StatusBadRequest)
				}
			}))
			defer server.Close()
			baseURL, err := url.Parse(server.URL + "/")
			require.NoError(t, err)
			service := NewGithubServiceImpl(server.Client(), baseURL)
			projects, err := service.FetchRepositories(context.Background())
			if failNextPage {
				require.ErrorContains(t, err, "page 2")
				assert.Nil(t, projects)
				return
			}
			require.NoError(t, err)
			require.Len(t, projects, 101)
			for i, project := range projects {
				assert.Equal(t, int64(i+1), project.Id)
			}
		})
	}
}

func TestFetchRepositoriesPreservesContextErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("canceled request should not reach the server")
	}))
	defer server.Close()
	baseURL, err := url.Parse(server.URL)
	require.NoError(t, err)
	service := NewGithubServiceImpl(server.Client(), baseURL)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	projects, err := service.FetchRepositories(ctx)
	require.ErrorIs(t, err, context.Canceled)
	assert.Nil(t, projects)
}
