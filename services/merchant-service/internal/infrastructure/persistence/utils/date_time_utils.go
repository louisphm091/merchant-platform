package utils

import "time"

func OffsetDateTimeNow() string {
	return time.Now().Format("2006-01-02T15:04:05.000-07:00")
}
