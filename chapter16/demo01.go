package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test()  hello world!" + strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}
}
func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("hello Gogang!", i+1)
		time.Sleep(time.Second)
	}
}
