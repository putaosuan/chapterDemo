package main

import "fmt"

//结构体可以使用嵌套匿名结构体所有的字段和方法，即首字母大写或者小写的字段、方法，都可以使用
type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk ", a.Name)
}
func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
}

func (b *B) hello() {
	fmt.Println("B hello", b.Name)
}

func main() {
	var b B
	b.Name = "tom"
	b.age = 20
	b.SayOk()
	b.hello() //采用就近原则访问
	b.A.hello()
	fmt.Println(b)
}
