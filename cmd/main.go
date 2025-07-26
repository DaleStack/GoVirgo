package main

import (
	"fmt"
	"govirgo/internal/scaffolder"
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "create" || os.Args[2] != "app" {
		fmt.Println("Usage: go run cmd/main.go create app <app_name>")
		return
	}
	name := os.Args[3]
	err := scaffolder.CreateHandler(name)
	if err != nil {
		fmt.Println("Error creating handler:", err)
	}
}
