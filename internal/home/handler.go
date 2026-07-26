package home

import (
	"net/http"

	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common"
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common/templates"
)

type Handler struct {
	templateManager *templates.TemplateManager
}

func NewHandler(templateManager *templates.TemplateManager) *Handler {
	return &Handler{templateManager}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	err := h.templateManager.Render(w, "home.html", nil)

	if err != nil {
		common.HandleError(w, h.templateManager, err)
	}
}
