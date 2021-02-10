package goutils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenOrderID 生成订单号
// 基础做法 日期20191025时间戳1571987125435+3位随机数
func GenOrderID() string {
	date := time.Now().Format("20060102")
	timestamp := time.Now().UnixNano() / 1e6
	rn := rand.Intn(1000)

	return fmt.Sprintf("%s%d%03d", date, timestamp, rn)
}
