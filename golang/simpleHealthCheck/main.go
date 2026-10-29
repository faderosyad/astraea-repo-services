package main

import (
	"log"
	"net/http"
	"simpleHealthCheck/httpUtils"
)

func main() {
	http.HandleFunc("/", httpUtils.Server)
	http.HandleFunc("/healthz", httpUtils.HealthCheck)

	log.Println("Server running on :8080...")
	http.ListenAndServe(":8080", nil)
}
