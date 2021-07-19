package main

import "fmt"

func main() {
	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println(numArr01)

	var numArr02 = [3]int{3, 4, 5}
	fmt.Println(numArr02)

	var numArro3 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numArro3)

	var numArr04 = [...]int{1: 100, 0: 10, 2: 200}
	fmt.Println(numArr04)
}
