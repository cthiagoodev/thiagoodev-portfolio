package usecases

import (
	"reflect"
	"testing"
	"time"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/github"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/test/mocks"
	"go.uber.org/mock/gomock"
)

func TestSyncProjectsUseCaseImpl_Execute(t *testing.T) {
	t.Run("should fetch, parse, save and return projects", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		repository := mocks.NewMockProjectsRepository(ctrl)
		service := mocks.NewMockService(ctrl)

		now := time.Now()

		githubProjects := []github.Project{
			{
				Id:        123,
				Name:      "portfolio",
				HtmlUrl:   "https://github.com/cthiagoodev/portfolio",
				CreatedAt: now,
				UpdatedAt: now,
			},
		}

		expectedProjects := []entities.Project{
			{
				ExternalId: "123",
				Name:       "portfolio",
				Url:        "https://github.com/cthiagoodev/portfolio",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
		}

		service.
			EXPECT().
			FetchRepositories().
			Return(githubProjects, nil)

		repository.
			EXPECT().
			CreateOrUpdate(expectedProjects[0]).
			Return(expectedProjects[0], nil)

		useCase := NewSyncProjectsUseCaseImpl(repository, service)

		result, err := useCase.Execute()

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(result) != len(expectedProjects) {
			t.Fatalf(
				"expected %d projects, got %d",
				len(expectedProjects),
				len(result),
			)
		}

		if reflect.DeepEqual(result, expectedProjects) == false {
			t.Errorf(
				"expected project %+v, got %+v",
				expectedProjects[0],
				result[0],
			)
		}
	})
}
