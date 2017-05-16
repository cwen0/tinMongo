package utils

import (
	"fmt"
	"time"
)

func Date(t time.Time) string {
	return fmt.Sprintf("%v-%v-%v-%v-%v-%v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Month(), t.Second())
}
