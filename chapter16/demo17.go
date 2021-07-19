package main

import "fmt"

func main() {
	//使用select可以解决从管道取数据的堵塞问题
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//传统方法在遍历管道时，如果不关闭阻塞而导致deadlock
	//问题在实际开发中，可能我们不好确定什么时候关闭管道
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从intChan读取的数据%s\n", v)
		default:
			fmt.Println("都读取不到，不玩了")
			return
		}
	}
}
