package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Problem 4 — HTTP Server (Basic Web)

Goal:

	Start an HTTP server on :8080 that responds to GET /hello with "Hello from Go!".
	If provided ?name=Mike, respond "Hello, Mike!".

Why:

	Teaches net/http, routing basics, and writing handlers.

Run:

	go run ./4_http_server
	curl 'http://localhost:8080/hello'
	curl 'http://localhost:8080/hello?name=Mike'

TODOs:
 1. Implement an http.HandlerFunc that reads `name` from the querystring.
 2. If name is empty, default to "from Go".
 3. Start the server on :8080 and log fatal on error.
*/
func main() {
	http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	name := q.Get("name")
	if name == "" {
		name = "from go"
	}
	fmt.Fprintf(w, "Hello %s", name)
}
