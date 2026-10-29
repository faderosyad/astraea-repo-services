package httpUtils

import (
	"fmt"
	"net/http"
	"time"
)

func Server(w http.ResponseWriter, req *http.Request) {
	//Return specific string to user - Server
	fmt.Fprint(w, "Welcome\n\n")
	fmt.Println()

	//Return specific time to user - Server
	time := time.Now()
	fmt.Fprintf(w, "%v\n\n", time.Format("2006-01-02 15:04:05"))
}