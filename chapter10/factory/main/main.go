package main

import (
	"chapterDemo/chapter10/factory/model"
	"fmt"
)

func main() {
	//var stu=model.Student{
	//	Name:"tom",
	//	Score: 78.9,
	//}
	var stu = model.NewStudent("tom", 67.8)
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
