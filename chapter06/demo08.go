package main

import "fmt"

func sum(n int, n2 ...int) int {
	sum := 1
	for i := 0; i < len(n2); i++ {
		sum += n2[i]
	}
	return sum
}
func main() {
	num := sum(1, 2, 3, 4, 5)
	fmt.Println(num)
}
