package main

import "fmt"

/*
	执行顺序：全局变量定义》init函数》main函数
*/
var age = test1()

func test1() int {
	fmt.Println("test1()...")
	return 0
}
func init() {
	fmt.Println("init()...")
}
func main() {
	fmt.Println("main()...")
}
