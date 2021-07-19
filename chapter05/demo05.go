package main

import "fmt"

func main() {
	/*
		打印1-100之间所有是9的倍数的整数的个数及总和
	*/
	var sum int
	var count int
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Println("count=", count, "sum=", sum)
	fmt.Println("-------------------------")
	var n int
	fmt.Println("请输入n")
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		fmt.Println(i, "+", n-i, "=", n)
	}
}
