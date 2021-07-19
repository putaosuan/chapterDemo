package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//求出一个数组的最大值，并得到对应的下标
	var intArr [6]int = [...]int{1, -1, 9, 90, 11, 9000}
	maxval := intArr[0]
	index := 0
	for i := 0; i < len(intArr); i++ {
		if maxval < intArr[i] {
			maxval = intArr[i]
			index = i
		}
	}
	fmt.Println(maxval, index)
	//求出一个数组的和和平均值
	//随机生成五个数，并将其反转打印，复杂应用
	var intArr3 [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(10) + 1
	}
	for i := 0; i < len(intArr3); i++ {
		fmt.Print(intArr3[i], " ")
	}
	fmt.Println()
	fmt.Println("转换后")
	for i := 0; i < len(intArr3)/2; i++ {
		t := intArr3[i]
		intArr3[i] = intArr3[len(intArr3)-i-1]
		intArr3[len(intArr3)-i-1] = t
	}
	for i := 0; i < len(intArr3); i++ {
		fmt.Print(intArr3[i], " ")
	}
}
