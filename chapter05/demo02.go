package main

import "fmt"

func main() {
	/*
		请编写一个程序，该程序可以接受一个字符，如a,b,c.. ,a表示星期一，b表示星期二，以此类推，根据用户
		的输入显示相应的信息
	*/
	var key byte
	fmt.Println("请输入一个字符，如a,b,c等")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期日")
	default:
		fmt.Println("都不是")
	}
}
