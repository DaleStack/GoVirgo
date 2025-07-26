package note

import (
	"govirgo/internal/router"
	"net/http"
)

func RegisterRoutes(router *router.Router) {
	router.Handle(http.MethodGet, "/note", getHandler)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("note handler is alive!"))
}
