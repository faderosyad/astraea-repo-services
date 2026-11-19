package main

import (
	"log"
	"net/http"
	"simpleHealthCheck/httpUtils"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Prevent running twice when browser requests /favicon.ico
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}

		// Run both Server and HealthCheck together
		httpUtils.Server(w, req)
		httpUtils.HealthCheck(w, req)
	})

	log.Println("Server running on :8080...")
	http.ListenAndServe(":8080", nil)
}
