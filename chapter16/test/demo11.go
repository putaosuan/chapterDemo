package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name     string
	Age      int
	Addresss string
}

func main() {
	var perChan chan Person
	perChan = make(chan Person, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10; i++ {
		person := Person{}
		person.Name = "hero" + fmt.Sprintf("%d", rand.Intn(100))
		person.Age = rand.Intn(20)
		person.Addresss = "北京"
		perChan <- person
	}
	close(perChan)
	for v := range perChan {
		fmt.Println(v)
	}

}
