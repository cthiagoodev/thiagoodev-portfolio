package about

import (
	"encoding/json"
	"errors"
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

	if errors.Is(err, common.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(about)
}
