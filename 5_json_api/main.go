package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
Problem 5 — JSON API (Encoding/Decoding)

Goal:
  Serve GET /users and return JSON array of users with fields: id, name, email.
  Bonus: Support POST /users to add a user (read JSON from body).

Why:
  Teaches structs, JSON tags, encoding/decoding, Content-Type.

Run:
  go run ./5_json_api
  curl -s 'http://localhost:8080/users' | jq

TODOs:
  1) Define a User struct with JSON tags.
  2) Hold an in-memory slice of users.
  3) Implement GET /users to write JSON and set Content-Type.
  4) (Bonus) Implement POST /users to append to the slice.
*/
func main() {
	// TODO: implement
	_ = json.Marshal
	_ = http.HandleFunc
	log.Println("TODO: start server on :8080")
}
