package conectivity

import (
	"context"
	"github.com/api_base/internal/domain/model"
	"github.com/go-chi/chi"
	"net/http"
)

type HandlerFunc interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type Service interface {
	Get(ctx context.Context, id string) (*model.Model, error)
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
		_, _ = w.Write([]byte("invalid_id"))
	}
	_, err := h.service.Get(ctx, id)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}
	_, _ = w.Write([]byte("ok"))
}
