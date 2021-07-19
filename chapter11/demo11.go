package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	b = a.(Point)
	fmt.Println(b)

	var x interface{}
	var b2 float32 = 3.1
	x = b2

	//类型断言（带检测）
	if y, ok := x.(float64); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T，值是%v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行")
}
