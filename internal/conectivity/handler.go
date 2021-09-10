package conectivity

import (
	"context"
	"github.com/api_base/internal/conectivity/web"
	"github.com/api_base/internal/domain/model"
	"github.com/go-chi/chi"
	"net/http"
)

type HandlerFunc interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type Service interface {
	Get(ctx context.Context, id string) (*model.User, error)
}

type handler struct {
	service Service
}

func NewHandlerFunc(srv Service) HandlerFunc {
	return &handler{service: srv}
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		_ = web.RespondJSON(w, []byte("invalid_id"), http.StatusBadRequest)
	}
	user, err := h.service.Get(ctx, id)
	if err != nil {
		_ = web.RespondJSON(w, err, http.StatusInternalServerError)
	}
	_ = web.RespondJSON(w, user, http.StatusOK)
}
