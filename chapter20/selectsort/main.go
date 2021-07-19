package main

import (
	"fmt"
	"math/rand"
	"time"
)

//选择排序
//编写函数
func SelectSort(arr *[80000]int) {

	for i := 0; i < len(arr)-1; i++ {
		max := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
				maxIndex = j
			}
		}
		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
	}
}
func main() {
	//定义一个数组
	//var arr = [5]int{20,34,19,100,80}
	//fmt.Println(arr)
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(90000)
	}
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Println(end - start)
	//fmt.Println(arr)

}
