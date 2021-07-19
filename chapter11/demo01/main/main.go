package main

import (
	"chapterDemo/chapter11/demo01/model"
	"fmt"
)

func main() {
	p := model.NewPerson("小米")
	p.SetAge(10)
	p.SetSal(15000)
	fmt.Println(p)
	fmt.Println(p.Name, "age=", p.GetAge(), "sal=", p.GetSal())
}
