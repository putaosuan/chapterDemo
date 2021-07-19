package main

import (
	"fmt"
)

//二分查找
func BinaryFind(arr []int, leftIndex int, rightIndex int, FindVal int) {
	if leftIndex > rightIndex {
		fmt.Println("没有找到")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if arr[middle] > FindVal {
		BinaryFind(arr, leftIndex, middle-1, FindVal)
	} else if arr[middle] < FindVal {
		BinaryFind(arr, middle+1, rightIndex, FindVal)
	} else {
		fmt.Println("zhaodaol")
	}
}
func main() {
	arr := []int{1, 8, 10, 89, 1000, 1234}
	//判断
	BinaryFind(arr, 0, len(arr)-1, 1000)

}
