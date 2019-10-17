package example

import (
	"fmt"
	"time"
)

func init() {
	// 1. 获取时间戳
	// 1.1 单位为秒
	timeStampS := time.Now().Unix()
	// 1.2 单位为毫秒
	timeStampMs := time.Now().UnixNano() / (1000 * 1000)
	// 1.3 单位为微秒
	timeStampUs := time.Now().UnixNano() / 1000
	// 1.4 单位为纳秒
	timeStampNs := time.Now().UnixNano()

	fmt.Printf(" s: %d\nms: %d\nus: %d\nns: %d\n", timeStampS, timeStampMs, timeStampUs, timeStampNs)

	// 2. 时间戳转换为日期字符串
	date1 := time.Now().Format("2006-01-02 15:04:05")
	date2 := time.Now().Format("2006/01/02 15:04:05")

	fmt.Printf("date1: %s\ndate2: %s\n", date1, date2)

	// 3. 日期字符串转换为时间戳

}
