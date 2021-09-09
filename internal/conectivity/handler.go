package conectivity

import "github.com/go-chi/chi"

type RouterHandler interface {
	Handler() *chi.Mux
}

type routerHandler struct {
	handler HandlerFunc
}

func (rh routerHandler) Handler() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/get/{id}", rh.handler.Get)
	return r
}

func NewRouterHandler(handler HandlerFunc) RouterHandler {
	return &routerHandler{
		handler: handler,
	}
}
