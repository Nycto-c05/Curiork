package main

import (
	"log"
	"minio-go-s3/internal/repository"
	object "minio-go-s3/internal/storage"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	repo   repository.MetaRepository
	object object.ObjectStore
}

type config struct {
	addr string
	db   dbConfig
	blob blobConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type blobConfig struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	bucket          string
	useSSL          bool
}

// ---- application methods
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Println("Server Running on", app.config.addr)
	return srv.ListenAndServe()
}
