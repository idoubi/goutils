package examples

import (
	"fmt"

	"github.com/idoubi/goutils"
)

func ExampleGenUuid() {
	uuid := goutils.GenUuid()

	fmt.Println(uuid)

	// Output: ae8c3136-a86f-4635-9ff7-7899a5a31908
}

func ExampleGenSnowid() {
	snowid := goutils.GenSnowid()

	fmt.Println(snowid)

	// Output: 1577901599664115712
}
