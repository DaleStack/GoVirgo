package server

import (
	"govirgo/apps/task"
	"govirgo/internal/router"
	"log"
	"net/http"
)

func Start() {
	r := router.New()

	task.RegisterRoutes(r)

	addr := ":8080"
	log.Printf("GoVirgo listening on %s\n", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
