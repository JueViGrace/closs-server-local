package util

import (
	"fmt"
	"strconv"
	"time"
)

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func StringToFloat(s string, bitSize int) float64 {
	n, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		n = 0.0
	}

	return n
}

func FormatDateForResponse(d time.Time) string {
	return fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02d",
		d.Year(), d.Month(), d.Day(),
		d.Hour(), d.Minute(), d.Second(),
	)
}
