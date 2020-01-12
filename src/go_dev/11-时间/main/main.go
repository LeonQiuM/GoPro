package main

import (
	"fmt"
	"time"
)

//
/*
时间
*/
func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Date())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) // 纳秒

	// 时间戳转化为时间格式
	ret := time.Unix(1578816280, 0)
	// ret1:= time.Unix(1578816280,1578816436431695555)
	fmt.Println(ret.Year(), ret.Month(), ret.Day(), ret.Hour(), ret.Minute(), ret.Second())

	// Add
	// now + 1h
	fmt.Println(now.Add(time.Hour * 24).Hour())

	// 定时器
	//timer:= time.Tick(time.Second)
	//for t:= range timer{
	//	fmt.Println(t)  // 一秒钟执行一次
	//}

	// 时间格式化，与时间对象转换为字符串类型
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006-01-02_3:04:05"))
	fmt.Println(now.Format("2006-01-02_3:04:05.0000000"))
	fmt.Println(now.Format("2006-01-02_15:04:05"))
	fmt.Println(now.Format("2006-01-02_03:04:05 PM"))
	// 哪找格式字符串转换为时间戳
	timeObj, err := time.Parse("2006-01-02_15:04:05", "1992-05-27_12:33:10")
	if err != nil {
		fmt.Println("parse time failed")
	}
	fmt.Println(timeObj.Unix())
	fmt.Println(timeObj)

	// Sub,两个时间对象相减，返回时间间隔
	fmt.Println(now.Sub(timeObj).Hours())

	// sleep
	n := 10000000
	time.Sleep(100000)
	time.Sleep(time.Duration(n))
	//time.Sleep(1000*time.Millisecond)

	fmt.Println("-----------")
	f2()

}

/*
时区
*/
func f2() {
	now := time.Now() // 本地时间
	fmt.Println(now)
	// 明天的这个时间
	time.Parse("2006-01-02 15:04:05", "2020-01-12 19:30:05") // 解析出来为 utc 时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("failed")
		return
	}
	// 按照指定时区解析
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-01-12 19:30:05", loc)
	timeObj, err = time.ParseInLocation("2006-01-02 15:04:05", "2020-01-12 19:36:05", time.Local)
	if err != nil {
		fmt.Println("failed2")
	}
	fmt.Println(timeObj)
}
