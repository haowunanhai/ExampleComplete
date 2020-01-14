package example

import (
	"testing"
)

func TestExample(t *testing.T) {
	// 1. 获取当前时间戳
	getCurrentTimeStamp()

	// 2. 时间戳转换为日期字符串
	convertTimeStampToDateStr()

	// 3. 日期字符串转换为时间戳
	convertDateStrToTimeStamp()

	// 4. 获取当期的时间字符串（2019-10-26 23:18:00）
	getCurrentDateStr()
}
