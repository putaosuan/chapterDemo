package main

import "fmt"

func main() {
	//从键盘输入值
	//要求：可以从控制台接受用户信息：姓名，年龄，薪水，是否通过考试
	var name string
	var age byte
	var sal float64
	var isPass bool
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试：")
	fmt.Scanln(&isPass)
	fmt.Println(name, age, sal, isPass)
	//方法二：
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，用空格隔开")
	fmt.Scanf("%s %d %f %t\n", &name, &age, &sal, &isPass)
	fmt.Println(name, age, sal, isPass)

}
