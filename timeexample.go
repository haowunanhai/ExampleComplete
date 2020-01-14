package example

import (
	"fmt"
	"time"
)

// 1. 获取当前时间戳
func getCurrentTimeStamp() {
	// 1.1 单位为秒
	timeStampS := time.Now().Unix()
	// 1.2 单位为毫秒
	timeStampMs := time.Now().UnixNano() / (1000 * 1000)
	// 1.3 单位为微秒
	timeStampUs := time.Now().UnixNano() / 1000
	// 1.4 单位为纳秒
	timeStampNs := time.Now().UnixNano()

	fmt.Printf(" s: %d\nms: %d\nus: %d\nns: %d\n", timeStampS, timeStampMs, timeStampUs, timeStampNs)
}

// 2. 时间戳转换为日期字符串
func convertTimeStampToDateStr() {
	// 2.1 生成当前日期，时间
	date1 := time.Now().Format("2006-01-02 15:04:05")
	date2 := time.Now().Format("2006/01/02 15:04:05")
	date4 := time.Now().Format("2006-01-02")
	fmt.Printf("date1: %s\ndate2: %s\ndate4: %s\n", date1, date2, date4)
	// 2.2 将某个时间戳转成日期字符串
	timeStamp := 1572244341
	t := time.Unix(int64(timeStamp), 0)
	date3 := t.Format("20060102 15:04:05")

	fmt.Println("date3:", date3)
}

// 3. 日期字符串转换为时间戳
func convertDateStrToTimeStamp() {
	// 3.1 不考虑时区
	dateStr := "2019-10-26 23:18:00"
	t, _ := time.Parse("2006-01-02 15:04:05", dateStr)
	timeStamp := t.Unix()

	// 3.2 使用本地时区
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
	timeStampLocal := t1.Unix()

	fmt.Println("Convert dateStr:", dateStr, "to timeStamp:", timeStamp, "timeStamp local:", timeStampLocal)

	// 3.3 获取今日零点的时间戳
	t0, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	timeStamp0 := t0.Unix()

	fmt.Println("Today zero time:", timeStamp0)

}

// 4. 获取当期的时间字符串（2019-10-26 23:18:00）
func getCurrentDateStr() {
	timeStr := time.Now().String()
	fmt.Println("Time Now:", timeStr)

}
