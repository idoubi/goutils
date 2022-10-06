package examples

import (
	"fmt"

	"github.com/idoubi/goutils"
)

func ExampleMGetStr() {
	m := goutils.MixMap{
		"foo": "bar",
		"i":   12,
	}

	foo := goutils.MGetStr(m, "foo")
	i := goutils.MGetInt(m, "i")
	f := goutils.IsInMap(m, "foo")

	fmt.Println(foo, i, f)

	// Output: bar 12 true
}

func ExampleMultiline2Map() {
	multiline := "foo:bar\r\nnum: 123\r\nswitch: false"

	m := goutils.Multiline2Map(multiline)

	fmt.Println(m["num"])

	// Output: 123
}
