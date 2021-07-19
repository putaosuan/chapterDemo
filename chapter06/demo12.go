package main

import "fmt"

func AddUpper() func(int) int {
	var n = 10
	var str = "hello"
	return func(x int) int {
		n += x
		str += string(36)
		fmt.Println(str)
		return n
	}
}
func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
