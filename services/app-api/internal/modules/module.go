package modules

import (
	"github.com/cthiagoodev/portfolio/services/app-api/internal/modules/about"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/modules/social_contact"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	About         *about.AboutModule
	SocialContact *social_contact.SocialContactModule
}

func NewModule(pool *pgxpool.Pool) *Module {
	return &Module{
		About:         about.NewAboutModule(pool),
		SocialContact: social_contact.NewSocialContactModule(pool),
	}
}
