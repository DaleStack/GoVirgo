package server

import (
	"govirgo/apps/task"

	"github.com/go-chi/chi/v5"
)

func RegisterAllRoutes(r chi.Router) {
	//Register your routes here

	task.RegisterRoutes(r)

}
