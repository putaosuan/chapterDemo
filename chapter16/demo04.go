package main

import "fmt"

func main() {
	//1.创建一个可以存放3个int的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2.看看inChan是什么
	fmt.Printf("intChan的值=%p intChan本身的地址=%p\n", intChan, &intChan)
	//3.向管道写入数据
	intChan <- 20
	num := 211
	intChan <- num
	intChan <- 50
	//intChan <-40
	//4.看看管道的长度和容量
	fmt.Println(len(intChan), cap(intChan))
	//5.从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Println(len(intChan), cap(intChan))

	//6.在没有使用协程的情况下，如果我们的管道数据已经全部取出再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan
	//num5:=<-intChan
	fmt.Println("num3=", num3, "num4=", num4)

}
