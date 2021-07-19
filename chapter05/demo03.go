package main

import "fmt"

func main() {
	//使用switch把小写类型的char转换成大写
	var mychar byte
	fmt.Println("请输入字符")
	fmt.Scanf("%c", &mychar)
	switch mychar {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	default:
		fmt.Println("other")
	}

}
