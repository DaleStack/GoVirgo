package scaffolder

import (
	"errors"
	"os"
	"path/filepath"
	"text/template"
)

func CreateHandler(name string) error {
	dir := filepath.Join("apps", name)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	filePath := filepath.Join(dir, "routes.go")
	if _, err := os.Stat(filePath); err == nil {
		return errors.New("routes.go already exists — skipping scaffolding")
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	const handlerTemplate = `package {{.Name}}

import (
    "net/http"
    "govirgo/internal/core"
)

func RegisterRoutes(router *core.Router) {
    router.Handle(http.MethodGet, "/{{.Name}}", getHandler)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("{{.Name}} handler is alive!"))
}
`

	tmpl, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		return err
	}

	return tmpl.Execute(f, struct{ Name string }{Name: name})
}
