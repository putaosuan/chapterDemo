package main

import "fmt"

type Point struct {
	x int
	y int
}
type Rect struct {
	leftUp, rightDown Point
}
type Rect2 struct {
	leftUp, RightDown *Point
}

func main() {
	//结构体注意事项：
	/*
		结构体的所有字段在内存中是连续的
	*/
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("%p,%p,%p,%p\n", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	//
	r2 := Rect2{&Point{30, 40}, &Point{10, 20}}
	fmt.Printf("%p,%p\n", &r2.leftUp, &r2.RightDown)
	fmt.Printf("%p,%p\n", r2.leftUp, r2.RightDown)
}
