package main

import (
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", api.getUsersHandler) // the handler fn dont have to be part of a struct/impl of Handler interface
	mux.HandleFunc("POST /users", api.createUserHandler)

	srv := http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
		// srv.ErrorLog.Panic(err)
	}
}
