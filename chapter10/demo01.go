package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 2
	cat1.Color = "白色"
	cat1.Hobby = "吃。。。"
	fmt.Println("猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
	fmt.Println("hobby=", cat1.Hobby)
}
