package utils

import (
	"log"

	"github.com/google/uuid"
)

func GenUUID() uuid.UUID {

	uuid := uuid.New()
	return uuid
}

func IsUUID(input interface{}) bool {

	if val, ok := input.(uuid.UUID); ok {
		if val != uuid.Nil {
			log.Println("is a uuid and not nil")
			return true
		}
		return false
	} else {
		log.Println("is not a uuid")
		return false
	}
}
