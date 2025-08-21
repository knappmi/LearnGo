package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

/*
Problem 5 — JSON API (Encoding/Decoding)

Goal:

	Serve GET /users and return JSON array of users with fields: id, name, Email.
	Bonus: Support POST /users to add a user (read JSON from body).

Why:

	Teaches structs, JSON tags, encoding/decoding, Content-Type.

Run:

	go run ./5_json_api
	curl -s 'http://localhost:8080/users' | jq

TODOs:
 1. Define a User struct with JSON tags.
 2. Hold an in-memory slice of users.
 3. Implement GET /users to write JSON and set Content-Type.
 4. (Bonus) Implement POST /users to append to the slice.
*/

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users = []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
		{ID: 3, Name: "Sam Smith", Email: "sam@example.com"},
	}
	mu sync.RWMutex
)

func main() {
	http.HandleFunc("/users", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		defer mu.RUnlock()

		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, "failed to encode users: "+err.Error(), http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		// Read and validate request body
		dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)) // 1MB cap (defensive)
		dec.DisallowUnknownFields()
		var u User
		if err := dec.Decode(&u); err != nil {
			http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		if u.Name == "" || u.Email == "" {
			http.Error(w, "name and email are required", http.StatusBadRequest)
			return
		}

		// Assign an ID and append
		mu.Lock()
		maxID := 0
		for _, existing := range users {
			if existing.ID > maxID {
				maxID = existing.ID
			}
		}

		if u.ID == 0 {
			u.ID = maxID + 1
		}
		users = append(users, u)
		mu.Unlock()

		// Respond 201 with the created object (common REST pattern)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(u); err != nil {
			log.Printf("write response: %v", err)
		}

	default:
		// Tell clients what’s allowed
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
