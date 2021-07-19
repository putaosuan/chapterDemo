package main

import "fmt"

type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 0; i < 10; i++ {
		for i := 0; i < 8; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func (mu MethodUtils) Print2(m, n int) {
	for i := 1; i <= m; i++ {
		for i := 1; i <= n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main() {
	var mu MethodUtils
	mu.Print()
	mu.Print2(2, 3)
}
