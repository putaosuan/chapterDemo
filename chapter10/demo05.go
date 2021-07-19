package main

import "fmt"

//结构体是用户单独定义的类型，和其他类型转换时需要有完全相同的字段
type A struct {
	Num int
}
type B struct {
	Num int
}

func main() {
	var a A
	a.Num = 1
	var b B
	b = B(a)
	fmt.Println(b)
}
