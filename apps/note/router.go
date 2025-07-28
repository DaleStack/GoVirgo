package note

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
    r.Route("/note", func(r chi.Router) {
        r.Get("/", getAllHandler)
        r.Get("/{id}", getByIDHandler)
        r.Post("/", createHandler)
        r.Delete("/{id}", deleteHandler)
    })
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    id := chi.URLParam(r, "id")
    w.Write([]byte("GET note with ID: " + id))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("POST new note"))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    id := chi.URLParam(r, "id")
    w.Write([]byte("DELETE note with ID: " + id))
}
