package utils

import (
	"strings"
)

func StringSplit(s string, sep string) []string {
    ss := strings.Split(s, sep)
    return ss
}
