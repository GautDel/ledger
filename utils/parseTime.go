package utils

import (
	"log"
	"time"
)

func ParseTime(layout, str string) time.Time {
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println("Error parsing time:", err)
		return time.Time{}
	}
    t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
    
	return t
}
