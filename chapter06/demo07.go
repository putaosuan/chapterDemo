package main

import "fmt"

func main() {
	//自定义数据类型 别名
	type myInt int
	var mi myInt = 60
	var i int
	i = int(mi)
	fmt.Println(i)
	fmt.Printf("%T,%T\n", mi, i)
}
