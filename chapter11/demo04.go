package main

import "fmt"

type A4 struct {
	Name string
	Age  int
}
type B4 struct {
	Name string
	Age  int
}
type C struct {
	A4
	B4
}

//如果struct嵌套了一个有名的结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或者方法时，
//必须带上结构体的名字
type D struct {
	a A4
}

func main() {
	var c C
	c.A4.Name = "xiao"
	fmt.Println(c.A4.Name)

	var d D
	d.a.Name = "nihao"
	fmt.Println(d)
}
