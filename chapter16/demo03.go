package main

import (
	"fmt"
	"sync"
	"time"
)

//计算1-120各个数的阶乘，并放入map中
var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex //是一个全局互斥锁 sync是包：全称是synchornized 同步
)

func test3(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	for i := 1; i <= 20; i++ {
		go test3(i)
	}
	time.Sleep(5 * time.Second)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("myMap[%v]=%v\n", i, v)
	}
	lock.Unlock()
}
