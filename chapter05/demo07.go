package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成随机数
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))

	//lable:
	for i := 1; i <= 4; i++ {
	lable1:
		for j := 1; i <= 10; j++ {
			if j == 3 {
				break lable1
			}
			fmt.Println("j=", j)
		}
	}
}
