package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}
func test18() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test18()发生错误", err)
		}
	}()
	var myMap map[string]int
	myMap["1"] = 5
}
func main() {
	go sayHello()
	go test18()
	for i := 0; i < 10; i++ {
		fmt.Println("main()ok=", i)
		time.Sleep(time.Second)
	}
}
