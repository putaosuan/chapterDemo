package main

import (
	"fmt"
	"time"
)

//要求统计1-20000中，哪些是素数

func putNum(intChan chan int) {
	for i := 1; i <= 20000; i++ {
		intChan <- i
	}
	close(intChan)
}
func primeNum(intChan chan int, resChan chan int, exitChan chan bool) {

	for {
		flag := true
		//time.Sleep(10*time.Millisecond)
		num, ok := <-intChan
		if !ok {
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明不是素数
				flag = false
				break
			}
		}
		if flag {
			resChan <- num
		}
	}
	fmt.Println("有一个primeNum因为取不到数据，退出了")
	exitChan <- true
}
func main() {
	intChan := make(chan int, 100)
	resChan := make(chan int, 20000)
	exitChan := make(chan bool, 4)
	start := time.Now().Unix()
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, resChan, exitChan)
	}
	//go func() {
	//
	//}()
	for {
		if len(exitChan) == 4 {
			end := time.Now().Unix()
			fmt.Println("耗费的时间是", start-end)
			close(resChan)
			fmt.Println("数量是：", len(resChan))
			break
		}
	}

	for v := range resChan {
		fmt.Println(v)
	}
	fmt.Println("main结束")
}
