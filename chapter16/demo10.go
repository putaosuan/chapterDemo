package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i <= 10; i++ {
		intChan <- i
		fmt.Println("writeData i", i)
		//time.Sleep(time.Second)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("读取成功，读到的数据是：%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
