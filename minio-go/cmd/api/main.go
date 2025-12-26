package main

import (
	"log"
	"minio-go-s3/internal/blob"
	"minio-go-s3/internal/db"
	"minio-go-s3/internal/repository"
	object "minio-go-s3/internal/storage"
)

func main() {

	config := &config{
		addr: ":8083",
		db: dbConfig{
			addr:         "postgresql://postgres:postgres@localhost:5432/paste?sslmode=disable",
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15m",
		},
		blob: blobConfig{
			endpoint:        "localhost:9000",
			accessKeyID:     "nycto",
			secretAccessKey: "nycto1234",
			bucket:          "pastes",
			useSSL:          false,
		},
	}

	//connect to db, verify conn, and get client
	db, err := db.NewDb(config.db.addr, config.db.maxOpenConns, config.db.maxIdleConns, config.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Print("DB connection pool established")

	//connect to s3
	s3Client, err := blob.NewObj(config.blob.endpoint, config.blob.accessKeyID, config.blob.secretAccessKey, config.blob.useSSL)
	if err != nil {
		log.Panic(err)
	}
	log.Print("Object Storage Connnection established")

	//Initialize impl of repo and object stores
	repo := repository.NewPostgresMetaRepository(db)
	objectStore := object.NewMinioStore(s3Client, config.blob.bucket)

	//Create application
	app := &application{
		config: *config,
		repo:   repo,
		object: objectStore,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
