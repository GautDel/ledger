package utils

import (
	"strconv"
	"time"
)

func GenInvoiceID(count int) string {
	curr := time.Now()
	y := strconv.Itoa(curr.Year())
	m := zeroPadding(int(curr.Month()))
	c := zeroPadding(count + 1)

	str := y + m + c

	return str
}

func zeroPadding(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}