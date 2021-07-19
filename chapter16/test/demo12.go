package main

import "fmt"

////计算1-120各个数的阶乘，并放入map中
var (
	myMap = make(map[int]int, 10)
)

func test12(n int, mapChan chan int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	mapChan <- res

}
func main() {
	mapChan := make(chan int, 10)
	for i := 1; i <= 12; i++ {
		go test12(i, mapChan)
		myMap[i] = <-mapChan
	}
	fmt.Println(myMap)

}
