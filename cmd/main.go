package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"ledgerbolt.systems/internal/handlers"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	mux := http.NewServeMux()

	handlers.ApiRouter(mux)

	port := os.Getenv("PORT")
	fmt.Printf("Server started on PORT: %s...", port)
    servErr := http.ListenAndServe(":" + port, mux)

	if servErr != nil {
		log.Fatal(500, servErr)
	}
}
