package httpUtils

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	//Return Google status - Client
	sites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://x.com",
		"https://linkedin.com",
	}

	for i, v := range sites {
		resp, err := http.Get(v)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Fprintf(w, "%d Hit Response Code: %v\n", i, resp.Status)
	}
}
