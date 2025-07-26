package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	*chi.Mux
}

func New() *Router {
	return &Router{chi.NewRouter()}
}

func (r *Router) MethodHandle(method, pattern string, handler http.HandlerFunc) {
	r.Method(method, pattern, handler)
}
