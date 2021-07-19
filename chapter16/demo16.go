package main

func main() {
	//只写
	var chan1 chan<- int
	chan1 = make(chan int, 10)
	chan1 <- 1
	var chan2 <-chan int
	//只读
	chan2 = make(chan int, 10)
	<-chan2

}
