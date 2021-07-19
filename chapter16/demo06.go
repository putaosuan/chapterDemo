package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var catChan chan Cat
	catChan = make(chan Cat, 3)
	cat1 := Cat{
		"tom",
		2,
	}
	cat2 := Cat{
		"niaho",
		5,
	}
	catChan <- cat1
	catChan <- cat2

	cat11 := <-catChan
	cat12 := <-catChan
	fmt.Println(cat11, cat12)
}
