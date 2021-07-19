package main

import "fmt"

func main() {
	/*
		有两个变量a和b，要求对这两个变量进行交换，但不允许使用中间变量
	*/
	var a = 10
	var b = 20
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a=", a, "b=", b)
}
