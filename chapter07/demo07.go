package main

import "fmt"

func fbn(n int) []int {
	f := make([]int, n)
	f[0] = 1
	if n > 1 {
		f[1] = 1
	}

	for i := 2; i < n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}
func main() {
	fmt.Println(fbn(10))
}
