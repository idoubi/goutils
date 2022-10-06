package examples

import (
	"fmt"
	"time"

	"github.com/idoubi/goutils"
)

func ExampleTimestamp() {
	timestamp := goutils.Timestamp()
	fmt.Println(timestamp)

	// Output: 1665035157
}

func ExampleTimestampStr() {
	timestampStr := goutils.TimestampStr()
	fmt.Println(timestampStr)

	// Output: 1665035211
}

func ExampleTimeFormat() {
	now := time.Now()
	t1 := goutils.TimeFormat(now)
	t2 := goutils.TimeFormat(now, "2006/01/02 15:04:05")
	fmt.Printf("%s, %s", t1, t2)

	// Output: 2022-10-06 13:47:35, 2022/10/06 13:47:35
}

func ExampleStr2Timestamp() {
	str1 := "2022-10-06 13:47:35"
	timestamp1 := goutils.Str2Timestamp(str1)

	str2 := "2018-06-08T10:34:56+08:00"
	timestamp2 := goutils.Str2Timestamp(str2, time.RFC3339)

	str3 := "2022/10/06 13:47:35"
	timestamp3 := goutils.Str2Timestamp(str3, "2006/01/02 15:04:05")

	fmt.Println(timestamp1, timestamp2, timestamp3)

	// Output: 1665064055 1528425296 1665064055
}

func ExampleTimestamp2Str() {
	timestamp := time.Now().Unix()

	str1 := goutils.Timestamp2Str(timestamp)

	str2 := goutils.Timestamp2Str(timestamp, time.RFC3339)

	str3 := goutils.Timestamp2Str(timestamp, "2006/01/02 15:04:05")

	fmt.Println(str1, str2, str3)

	// Output: 2022-10-06 13:55:17 2022-10-06T13:55:17+08:00 2022/10/06 13:55:17
}

func ExampleTimeSeq() {
	seq := goutils.TimeSeq()

	fmt.Println(seq)

	// Output: 2022100613555741466782
}
