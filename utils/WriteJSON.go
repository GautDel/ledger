package utils

import (
	"encoding/json"
	"log"
)

func writeJSON(data interface{}) {
    
    encoderErr := json.NewEncoder(w).Encode(data)
    if encoderErr != nil {
        log.Println("Failed to encode JSON", encoderErr)
    }
}
