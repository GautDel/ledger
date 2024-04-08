package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/handlers"
	"ledgerbolt.systems/internal/middleware"
)

func main() {

	// LOAD ENV VARIABLES //
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// ESTABLISH DATABASE CONNECTION //
	db.Connect(os.Getenv("DB_URL"))

	// INITIALIZE ROUTES/SERVER //
	mux := http.NewServeMux()

	handlers.ApiRouter(mux)

	port := os.Getenv("PORT")

	fmt.Printf("Server started on PORT: %s...", port)

	err = http.ListenAndServe(":"+port, middleware.SetCommonHeaders(mux))
	if err != nil {
		log.Fatal(500, err)
	}
}
