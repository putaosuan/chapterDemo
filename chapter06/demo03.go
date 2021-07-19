package main

import "fmt"

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

//斐波那契函数
func main() {
	/*
		给一个整数，求出斐波那契数列
	*/
	var n int
	fmt.Println("请输入整数：")
	fmt.Scanln(&n)
	fmt.Println(n, "的斐波那契数是：", fib(n))

}
