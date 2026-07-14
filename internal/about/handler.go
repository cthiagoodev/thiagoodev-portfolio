package about

import (
	"encoding/json"
	"net/http"

	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common"
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common/templates"
)

type Handler struct {
	useCase         *UseCase
	templateManager *templates.TemplateManager
}

func NewHandler(useCase *UseCase, templateManager *templates.TemplateManager) *Handler {
	return &Handler{useCase, templateManager}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	about, err := h.useCase.Get(r.Context())

	if err != nil {
		common.HandleError(w, h.templateManager, err)
		return
	}

	json.NewEncoder(w).Encode(&about)
}
