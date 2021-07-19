package main

import "fmt"

//GO中，函数也是一种数据类型，可以赋值给一个变量
func getsum(n1 int, n2 int) int {
	return n1 + n2
}
func main() {
	a := getsum
	fmt.Println(a(1, 2))
}
