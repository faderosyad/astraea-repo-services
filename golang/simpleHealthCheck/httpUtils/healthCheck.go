package httpUtils

import (
	"log"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	sites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://x.com",
		"https://linkedin.com",
		"https://faderosyad.com",
	}

	for _, endpoint := range sites {

		resp, err := client.Get(endpoint)
		if err != nil {

			log.Printf("[HEALTHCHECK] Endpoint: %s | ERROR: %v", endpoint, err)
			continue
		}

		log.Printf("[HEALTHCHECK] Endpoint: %s | Status: %s | Code: %d", endpoint, resp.Status, resp.StatusCode)

		resp.Body.Close()
	}
}
