package main

import "fmt"

//如果一个方式实现了String()这个方法，那么fmt.println默认会调用这个变量的String()进行输出
type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v],Age=[%v]", stu.Name, stu.Age)
	return str
}
func main() {
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	//如果实现了*Student类型的String()方法，就会自动调用
	fmt.Println(&stu)
}
