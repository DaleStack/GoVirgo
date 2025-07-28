package task

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var mockTasks = []map[string]string{
	{"id": "1", "title": "Buy groceries"},
	{"id": "2", "title": "Write blog post"},
}

func RegisterRoutes(r chi.Router) {
	r.Route("/task", func(r chi.Router) {
		r.Get("/", getAllHandler)
		r.Get("/{id}", getByIDHandler)
		r.Post("/", createHandler)
		r.Delete("/{id}", deleteHandler)
	})
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockTasks)
}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	for _, task := range mockTasks {
		if task["id"] == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Task not Found", http.StatusNotFound)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST new task"))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("DELETE task with ID: " + id))
}
