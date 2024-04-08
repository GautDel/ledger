package handlers

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("Welcome to the home page!"))

	if err != nil {
		log.Println("Error writing response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
