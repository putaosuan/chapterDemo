package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println(n)
	}

}
func main() {
	test(5)
}
