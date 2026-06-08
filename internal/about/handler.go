package about

import (
	"encoding/json"
	"net/http"

	"github.com/cthiagoodev/thiagoodev-portfolio/server/internal/common"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{useCase}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	about, err := h.useCase.Get(r.Context())

	if err != nil {
		common.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(&about)
}
