package examples

import (
	"fmt"

	"github.com/idoubi/goutils"
)

func ExampleNonceStr() {
	nonce := goutils.NonceStr(16)
	fmt.Println(nonce)

	// Output: tIVEwZDqoYyGfzrm
}
