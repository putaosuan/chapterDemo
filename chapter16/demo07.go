package main

import "fmt"

type Cat7 struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	cat1 := Cat7{
		"nihao",
		5,
	}
	allChan <- cat1
	allChan <- 5
	allChan <- "jack"
	cat11 := <-allChan
	v1 := <-allChan
	v2 := <-allChan
	fmt.Println(cat11, v1, v2)
}
