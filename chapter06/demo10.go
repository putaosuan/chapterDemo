package main

import "fmt"

func main() {
	//匿名函数
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println(res1(1, 2))
	res2 := func(n1 int, n2 int) int {
		return n1 + n2
	}(5, 6)
	fmt.Println(res2)
}
