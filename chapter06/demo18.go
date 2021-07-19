package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//方法一：
	fmt.Printf("当前年月日 %d-%d-%d %d-%d-%d\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d-%d-%d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateStr)
	//方法二：
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006/01/02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

}
