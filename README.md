## GoVirgo
---
### Create Application
```bash
go run cmd/main.go create app <app_name>
```
this will scaffold an application inside **apps/**

### Register App
navigate to **server/app_registry.go** and register the app route
```Go
package server

import (
	"govirgo/apps/app_name"

	"github.com/go-chi/chi/v5"
)

func RegisterAllRoutes(r chi.Router) {
	//Register your routes here

	app_name.RegisterRoutes(r)

}
```