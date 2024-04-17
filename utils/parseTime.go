package utils

import (
	"time"
)

func ParseTime(layout string, str time.Time) time.Time {

    t := time.Date(str.Year(), str.Month(), str.Day(), 0, 0, 0, 0, time.UTC)

	return t
}


