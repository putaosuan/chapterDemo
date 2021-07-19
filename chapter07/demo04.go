package main

import "fmt"

//切片的copy操作
func main() {
	var slice []int = []int{1, 2, 3, 4, 5}
	var slice2 = make([]int, 2)
	copy(slice2, slice)
	fmt.Println(slice2)
}
