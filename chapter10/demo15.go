package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func main() {
	var stu1 = Stu{"小米", 10}
	stu2 := Stu{"小米", 10}
	var stu3 = Stu{
		Name: "xiaomi",
		Age:  10,
	}
	stu4 := Stu{
		"xiaomi",
		10,
	}
	fmt.Println(stu1, stu2, stu3, stu4)
	var stu5 *Stu = &Stu{"xiaomi", 10}
	stu6 := &Stu{"xiaomi", 10}
	fmt.Println(*stu5, *stu6)
}
