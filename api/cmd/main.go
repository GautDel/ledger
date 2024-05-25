package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

	rtr := handlers.New()
    
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	certFile := "localhost.crt" // Update with your SSL certificate file path
	keyFile := "localhost.key"   // Update with your SSL private key file path

	log.Print("Server listening on https://" + host + ":" + port)
	// Start the HTTPS server
	if err := http.ListenAndServeTLS(host+":"+port, certFile, keyFile, rtr); err != nil {
		log.Fatalf("There was an error with the https server: %v", err)
	}
	
}
