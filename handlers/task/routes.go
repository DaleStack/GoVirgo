package task

import (
	"govirgo/internal/router"
	"net/http"
)

func RegisterRoutes(router *router.Router) {
	router.Handle(http.MethodGet, "/task", getHandler)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("task handler is alive!"))
}
