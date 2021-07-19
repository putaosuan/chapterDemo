package main

import (
	"fmt"
	"math/rand"
	"time"
)

//插入排序
func InsertSort(arr *[80000]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		//fmt.Println(arr)
	}

}
func main() {
	//arr:=[5]int{23,0,12,56,34}
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(90000)
	}
	start := time.Now().Unix()
	//fmt.Println(arr)
	InsertSort(&arr)
	end := time.Now().Unix()
	fmt.Println(end - start)
}
