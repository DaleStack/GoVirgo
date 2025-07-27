package server

import (
	"govirgo/apps/task"
	"govirgo/internal/router"
	"log"
	"net/http"
)

func BuildHandler() http.Handler {
	r := router.New()

	//Register application route here
	task.RegisterRoutes(r)

	//*IGNORE: Test Route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return r
}

func Start() {
	handler := BuildHandler()

	addr := ":8080"
	log.Printf("GoVirgo listening on 127.0.0.1%s\n", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
