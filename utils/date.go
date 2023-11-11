package utils

import (
	"time"
)

func GetCurrentTime() string {
	t := time.Now().Format("01-01-2006 3:04pm")
	return t
}
