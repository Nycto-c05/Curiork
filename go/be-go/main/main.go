package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User"))
}

func (s *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created User"))
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	api := &api{addr: ":8080"}

	//for routing
	mux := http.NewServeMux() //implemetns handler fn interface

	//can configure shit in this struct below for the http server
	srv := &http.Server{
		Addr: api.addr,
		// Handler: api,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler) // the handler dont have to be part of a struct/impl of Handler interface
	mux.HandleFunc("POST /users", api.createUserHandler)

	log.Fatal(srv.ListenAndServe()) // instead of err := fn() if err!=nil
}
