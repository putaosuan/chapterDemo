package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond * 10)
		if i == 100 {
			break
		}
	}
}
