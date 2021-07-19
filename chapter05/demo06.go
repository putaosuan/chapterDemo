package main

import "fmt"

func main() {
	/*
		打印金字塔
	*/
	var n int
	fmt.Println("请输入金字塔的层数")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || i == n || j == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
