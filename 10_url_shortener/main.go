package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
Problem 10 — Mini Project: URL Shortener (Maps, HTTP, JSON)

Goal:
  Build a simple in-memory URL shortener with endpoints:
    POST /shorten  -> body: {"url":"https://example.com"} returns {"short":"abc123"}
    GET  /<short>  -> redirects (302) to the original URL

Why:
  Teaches maps, HTTP handlers, JSON parsing, and basic routing.
  Optional: persistence to file; add TTL or stats.

Run:
  go run ./10_url_shortener
  curl -XPOST -d '{"url":"https://example.com"}' -H 'Content-Type: application/json' localhost:8080/shorten -i

TODOs:
  1) Create a map[string]string for short->original.
  2) Implement POST /shorten to parse JSON and generate a short code.
  3) Implement GET /{short} to look up and redirect.
  4) Start the server on :8080.
*/
func main() {
	// TODO: implement
	_ = json.Marshal
	_ = http.HandleFunc
	log.Println("TODO: start server on :8080")
}
