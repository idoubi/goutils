package examples

import (
	"fmt"

	"github.com/idoubi/goutils"
)

func ExampleIsIntInArr() {
	intArr := []int{1, 4, 3, 12}
	var i int = 3
	res1 := goutils.IsIntInArr(intArr, i)

	strArr := []string{"foo", "bar", "test"}
	str := "bar2"
	res2 := goutils.IsStrInArr(strArr, str)

	fmt.Println(res1, res2)

	// Output: true false
}
