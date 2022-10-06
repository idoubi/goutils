package goutils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const defaultTimeFormatter = "2006-01-02 15:04:05"

// Timestamp: Get current timestamp
func Timestamp() int64 {
	return time.Now().Unix()
}

// TimestampStr: Get current timestamp with format string
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

// Str2Timestamp: format timestr to unix timestamp
func Str2Timestamp(str string, format ...string) int64 {
	f := defaultTimeFormatter
	if len(format) > 0 {
		f = format[0]
	}

	t, _ := time.Parse(f, str)

	return t.Unix()
}

// Timestamp2Str: format unix timestamp to str
func Timestamp2Str(timestamp int64, format ...string) string {
	f := defaultTimeFormatter
	if len(format) > 0 {
		f = format[0]
	}

	t := time.Unix(timestamp, 0)

	return t.Format(f)
}

// TimeSeq: generate sequence number based on current timestamp
func TimeSeq() string {
	now := time.Now()
	nano := now.UnixNano()

	nows := now.Format("20060102150405")
	mill := nano % 1e6
	randNum := rand.Int63n(1000)

	return fmt.Sprintf("%s%d%d", nows, mill, randNum)
}
