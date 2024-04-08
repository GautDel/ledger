package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Info struct {
	Title string
	Desc  string
	Usage string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	info := Info{
		Title: "Ledger is an accounting API to be used by Aoife McNamara.",
		Desc:  "It will allow her to save clients in a database, as well as invoices, and will create PDF automatically.",
		Usage: "Docs are coming soon!",
	}

    encoderErr := json.NewEncoder(w).Encode(info) 
    if encoderErr != nil {
        log.Println("Failed to Encode JSON", encoderErr)
    }
}
