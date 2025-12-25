package main

import (
	"database/sql"
	"log"
	"minio-go-s3/internal/repository"
	object "minio-go-s3/internal/storage"

	"github.com/minio/minio-go/v7"
)

func main() {

	config := &config{
		addr: ":8080",
	}

	//Initialize impl of repo and object stores
	repo := repository.NewPostgresMetaRepository( /*db conn client*/ &sql.DB{})
	objectStore := object.NewMinioStore(&minio.Client{}, "bucketNamePlaceholder")

	app := &application{
		config: *config,
		repo:   repo,
		object: objectStore,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
