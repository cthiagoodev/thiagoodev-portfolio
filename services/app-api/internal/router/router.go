package router

import (
	"github.com/cthiagoodev/portfolio/services/app-api/internal/modules"
	"github.com/go-chi/chi/v5"
)

func NewRouter(applicationModule *modules.Module) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/about", applicationModule.About.Handler.Get)

	return chiRouter
}
