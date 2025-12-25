package main

import (
	"database/sql"
	"log"
	"minio-go-s3/internal/repository"
)

func main() {

	config := &config{
		addr: ":8080",
	}
	//Initialize impl of repo and object stores
	repo := repository.NewPostgresMetaRepository( /*db conn client*/ &sql.DB{})

	app := &application{
		config: *config,
		repo:   repo,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
