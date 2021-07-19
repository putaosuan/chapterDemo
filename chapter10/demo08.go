package main

import "fmt"

//Golang中的方法作用在指定的数据类型上（即：和指定的数据类型绑定），因此自定义类型都可以有方法，而不仅仅是struct
type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}
func (i *integer) change() {
	*i = *i + 1
}

//编写一个方法，可以改变i的值
func main() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println(i)
}
