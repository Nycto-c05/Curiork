package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

/*
	func (app *application) mount() *http.ServeMux {
			mux := http.NewServeMux()
			mux.HandleFunc("GET /v1/health", app.healthCheckerHandler)
			return mux
	}
*/
func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) //uses defer inside to recover

	// r.Get("/v1/health", app.healthCheckerHandler)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckerHandler)
	})
	return r
}

// since Handler takes impl of Handler interface, it means chi.Mux struct is impl of Handler, hence fn arg can be http.Handler instead, so can return type of mount() be
func (app *application) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute, //HTTP/1.1 keep-alive)
	}
	log.Printf("Sever running on %s", app.config.addr)

	return srv.ListenAndServe()
}
