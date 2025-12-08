package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var print = fmt.Println

type User struct {
	ID       string
	Name     string
	Products []string
}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /users/{id}", handleGetUserProducts)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	print("Listening on Port 8080")
	log.Fatal(server.ListenAndServe())

}

func handleGetUserProducts(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	user := User{
		ID:       userID,
		Name:     "Nigga singh",
		Products: []string{},
	}

	//req to db service to get orders of user
	// req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081/users/%s/products", userID), nil)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// r.Context(): request-scoped, cancels when client disconnects.
	// context.Background(): empty root context, never cancels unless you wrap it.
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://localhost:8081/users/%s/products", userID), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(data, &user.Products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, user)

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v) // marshalling
}

/*
Unmarshal: parse JSON from bytes; liek string or byte slice, data already in memory
 Decoder: parse JSON from a stream (io.Reader) without loading everything. thru a network givin in chunks
*/
