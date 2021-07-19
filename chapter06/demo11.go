package main

import "fmt"

//全局匿名函数
var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 - n2
	}
)

func main() {
	fmt.Println(Fun1(4, 2))
}
