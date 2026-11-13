package router

import (
	"github.com/cthiagoodev/portfolio/services/app-api/internal/modules"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(applicationModule *modules.Module) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.Logger)

	chiRouter.Get("/about", applicationModule.About.Handler.Get)
	chiRouter.Get("/social-contacts", applicationModule.SocialContact.Handler.GetAll)

	return chiRouter
}
