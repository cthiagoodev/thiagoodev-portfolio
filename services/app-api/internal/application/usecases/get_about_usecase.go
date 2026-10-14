package usecases

import (
	"context"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/repositories"
	domainusecases "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/usecases"
)

var _ domainusecases.GetAboutUseCase = (*GetAboutUseCaseImpl)(nil)

type GetAboutUseCaseImpl struct {
	aboutRepository repositories.AboutRepository
}

func NewGetAboutUseCase(aboutRepository repositories.AboutRepository) *GetAboutUseCaseImpl {
	return &GetAboutUseCaseImpl{
		aboutRepository: aboutRepository,
	}
}

func (u *GetAboutUseCaseImpl) Execute(ctx context.Context) (*entities.About, error) {
	return u.aboutRepository.Get(ctx)
}
