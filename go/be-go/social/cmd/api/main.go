package main

import (
	"log"

	"github.com/Nycto-c05/social/internal/env"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
