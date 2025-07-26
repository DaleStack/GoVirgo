package server

import (
	"govirgo/handlers/note"
	"govirgo/handlers/task"
	"govirgo/internal/router"
	"log"
	"net/http"
)

func Start() {
	router := router.NewRouter()

	task.RegisterRoutes(router)
	note.RegisterRoutes(router)

	addr := ":8080"
	log.Printf("🚀 GoVirgo listening on %s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
