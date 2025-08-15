package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
Problem 8 — REST API Client

Goal:
  Fetch posts from JSONPlaceholder and pretty-print the first 3 titles.
  Endpoint: https://jsonplaceholder.typicode.com/posts

Why:
  Teaches http.Get, reading response body, unmarshalling JSON into structs, and error handling.

Run:
  go run ./8_rest_api_client

TODOs:
  1) Perform an HTTP GET to the endpoint.
  2) Read the body and unmarshal into a struct (id, title, etc.).
  3) Print the first 3 titles with numbering.
*/
func main() {
	// TODO: implement
	_ = http.Get
	_ = io.ReadAll
	_ = json.Unmarshal
	fmt.Println("TODO")
}
