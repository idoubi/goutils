package goutils

import (
	"strconv"
	"time"
)

// Timestamp Get current timestamp
func Timestamp() int64 {
	return time.Now().Unix()
}

// TimestampStr Get current timestamp with format string
func TimestampStr() string {
	t := Timestamp()

	return strconv.Itoa(int(t))
}

// TimeFormat Format time to string
func TimeFormat(t time.Time, format ...string) string {
	f := "2006-01-02 15:04:05" // default format
	if len(format) > 0 {
		f = format[0]
	}

	return t.Format(f)
}
