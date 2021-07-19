package main

import "fmt"

func main() {
	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Print(arr3[i][j], " ")
		}
		fmt.Println()
	}
	for i, arr := range arr3 {
		for j, v := range arr {
			fmt.Printf("arr[%d][%d]=%v\t", i, j, v)
		}
		fmt.Println()
	}
}
