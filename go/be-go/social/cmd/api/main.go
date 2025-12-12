package main

import (
	"log"

	"github.com/Nycto-c05/social/internal/env"
	"github.com/Nycto-c05/social/internal/store"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
