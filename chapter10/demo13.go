package main

import "fmt"

type Student12 struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (stu Student12) say() string {
	infoStr := fmt.Sprintf("student的信息 name=%v,gender=%v,age=%v,id=%v,score=%v", stu.name,
		stu.gender, stu.age, stu.id, stu.score)
	return infoStr

}
func main() {
	var stu = Student12{
		"xiaoming",
		"男",
		22,
		123,
		66.7,
	}
	fmt.Println(stu.say())
}
