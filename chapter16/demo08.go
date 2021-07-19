package main

import "fmt"

type Cat8 struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 10)
	cat1 := Cat8{
		"nihao",
		5,
	}
	allChan <- 5
	allChan <- "jack"
	allChan <- cat1
	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("newCat=%T,newCat=%v\n", newCat, newCat)
	a := newCat.(Cat8)
	fmt.Println(a.Name)
}
