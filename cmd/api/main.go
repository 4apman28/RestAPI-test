package main

import (
	"fmt"
	"log"
	"net/http"
	"restapi-test/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)

	fmt.Println("Serve start in http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
