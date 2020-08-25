package examples

import (
	"fmt"
	"time"

	"github.com/idoubi/goutils"
)

func ExampleTimestamp() {
	t := goutils.Timestamp()
	fmt.Printf("%T, %v", t, t)

	// Output: int64, 1597381700
}

func ExampleTimeFormat() {
	now := time.Now()
	t1 := goutils.TimeFormat(now)
	t2 := goutils.TimeFormat(now, "2006/01/02 15:04:05")
	fmt.Printf("%s, %s", t1, t2)

	// Output: 2020-08-14 13:10:30, 2020/08/14 13:10:30
}
