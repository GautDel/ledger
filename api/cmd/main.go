package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"ledgerbolt.systems/internal/auth"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/handlers"
)

func main() {

	// LOAD ENV VARIABLES //
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// ESTABLISH DATABASE CONNECTION //
	db.Connect(os.Getenv("DB_URL"))

	auth, err := auth.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := handlers.New(auth)

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	log.Print("Server listening on http://" + host + ":" + port)
	if err := http.ListenAndServe("0.0.0.0:"+port, rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}

