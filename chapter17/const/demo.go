package main

import (
	"fmt"
)

func main() {
	var num int
	//常量生命的时候，必须赋值
	const tax int = 10
	fmt.Println(num, tax)

	//常量比较简洁的写法
	const (
		a = 1
		b = 2
	)
	fmt.Println(a, b)
	//还有一种比较专业的写法
	const (
		m    = iota
		n    = iota
		k, w = iota, iota
		u    = iota
	)
	fmt.Println(m, n, k, w, u)
	//Goland中没有常量名必须大写的规范
	//仍然通过首字母的大小写来控制常量的访问范围

}
