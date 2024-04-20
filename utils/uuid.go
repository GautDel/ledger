package utils

import (
	"github.com/google/uuid"
)

func GenUUID() uuid.UUID {

	uuid := uuid.New()
	return uuid
}

func IsUUID(input interface{}) bool {

	if val, ok := input.(uuid.UUID); ok {
		return val != uuid.Nil
	} else {
		return false
	}
}
