package main

import "fmt"

func sum1(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	n1++
	n2++
	n3 := n1 + n2
	fmt.Println("ok3 n3=", n3)
	return n3
}
func main() {
	res := sum1(10, 20)
	fmt.Println("res=", res)
}
