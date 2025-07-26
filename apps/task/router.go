package task

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
    r.Route("/task", func(r chi.Router) {
        r.Get("/", getAllHandler)
        r.Get("/{id}", getByIDHandler)
        r.Post("/", createHandler)
        r.Delete("/{id}", deleteHandler)
    })
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("GET all task"))
}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    w.Write([]byte("GET task with ID: " + id))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("POST new task"))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    w.Write([]byte("DELETE task with ID: " + id))
}
