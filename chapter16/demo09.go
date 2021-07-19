package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	//这里不能再写入数到channel
	//intChan<-300
	fmt.Println("okok...")
	n1 := <-intChan
	fmt.Println("n1=", n1)
	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		intChan2 <- i * 2
	}
	close(intChan2)
	//遍历管道不能用普通的for循环,在遍历时，如果channel没有关闭，则会出项deadlock的错误
	for v := range intChan2 {
		fmt.Println(v)
	}
}
