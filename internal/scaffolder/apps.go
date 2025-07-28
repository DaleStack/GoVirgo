package scaffolder

import (
	"errors"
	"os"
	"path/filepath"
	"text/template"
)

func CreateHandler(name string) error {
	dir := filepath.Join("apps", name)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	routerPath := filepath.Join(dir, "router.go")
	if _, err := os.Stat(routerPath); err == nil {
		return errors.New("router.go already exists — skipping scaffolding")
	}

	file, err := os.Create(routerPath)
	if err != nil {
		return err
	}
	defer file.Close()

	const routerTemplate = `package {{.Name}}

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
    r.Route("/{{.Name}}", func(r chi.Router) {
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
    w.Write([]byte("GET {{.Name}} with ID: " + id))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")s
    w.Write([]byte("POST new {{.Name}}"))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    id := chi.URLParam(r, "id")
    w.Write([]byte("DELETE {{.Name}} with ID: " + id))
}
`

	t := template.Must(template.New("router").Parse(routerTemplate))
	return t.Execute(file, struct{ Name string }{Name: name})
}
