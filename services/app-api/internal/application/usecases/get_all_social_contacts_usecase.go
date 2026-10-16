package usecases

import (
	"context"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/repositories"
	domainusecases "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/usecases"
)

var _ domainusecases.GetAllSocialContactsUseCase = (*GetAllSocialContactsUseCaseImpl)(nil)

type GetAllSocialContactsUseCaseImpl struct {
	socialContactRepository repositories.SocialContactRepository
}

func NewGetAllSocialContactsUseCase(socialContactRepository repositories.SocialContactRepository) *GetAllSocialContactsUseCaseImpl {
	return &GetAllSocialContactsUseCaseImpl{
		socialContactRepository: socialContactRepository,
	}
}

func (u *GetAllSocialContactsUseCaseImpl) Execute(ctx context.Context) ([]entities.SocialContact, error) {
	return u.socialContactRepository.GetAll(ctx)
}

