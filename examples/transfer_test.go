package examples

import (
	"fmt"

	"github.com/idoubi/goutils"
)

func ExampleTextEncode() {
	raw := "你好😝"
	new := goutils.TextEncode(raw)

	fmt.Println(new)

	// Output: 你好
}

func ExampleMap2XML() {
	m := map[string]string{
		"name":      "mike",
		"age":       "18",
		"isMale":    "true",
		"interests": `cdata:["basketball", "reading"]`,
	}

	x, _ := goutils.Map2XML(m, "xml")
	fmt.Printf("%s", x)

	// Output: <xml><name>mike</name><age>18</age><isMale>true</isMale><interests><![CDATA[["basketball", "reading"]]]></interests></xml>
}

func ExampleXML2JSON() {
	x := `<xml><return_code><![CDATA[SUCCESS]]></return_code>
	<return_msg><![CDATA[OK]]></return_msg>
	<appid><![CDATA[wx3497b8c1f5386a1e]]></appid>
	<mch_id><![CDATA[1498014222]]></mch_id>
	<nonce_str><![CDATA[jVwLeJZ6gsVgBYnm]]></nonce_str>
	<sign><![CDATA[8B6215C0CF3506214E29DEB787B578CD]]></sign>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<prepay_id><![CDATA[wx101208577821625cbca9eb7ea1fa7f0000]]></prepay_id>
	<trade_type><![CDATA[JSAPI]]></trade_type>
	</xml>`
	j, _ := goutils.XML2JSON([]byte(x))
	fmt.Printf("%s", j)

	// Output: {"xml": {"sign": "8B6215C0CF3506214E29DEB787B578CD", "result_code": "SUCCESS", "prepay_id": "wx101208577821625cbca9eb7ea1fa7f0000", "trade_type": "JSAPI", "appid": "wx3497b8c1f5386a1e", "return_msg": "OK", "mch_id": "1498014222", "nonce_str": "jVwLeJZ6gsVgBYnm", "return_code": "SUCCESS"}}
}
