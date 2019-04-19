package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	DumpHTTPRequest(r)
	fmt.Fprintf(w, "Hello from Go on Now 2.0!, use now logs to inspect dumped http request")
}

// DumpHTTPRequest for debugging
func DumpHTTPRequest(r *http.Request) {
	output, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Error dumping request:", err)
		return
	}
	log.Println(string(output))
}
