package modules

import (
	"github.com/cthiagoodev/portfolio/services/app-api/internal/modules/about"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	About *about.AboutModule
}

func NewModule(pool *pgxpool.Pool) *Module {
	return &Module{
		About: about.NewAboutModule(pool),
	}
}
