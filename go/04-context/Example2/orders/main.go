package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /users/{id}/products", handleGetUserProducts)
	// router.Handle() look into this

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Listening on port 8081")
	log.Fatal(server.ListenAndServe())

}

func handleGetUserProducts(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	log.Printf("Retrieving data from DB for userID : %s", userID)

	// db call
	time.Sleep(4 * time.Second)
	products := []string{"Prod1", "Prod2"}

	writeJSON(w, http.StatusOK, products)

}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v) // marshalling

}
