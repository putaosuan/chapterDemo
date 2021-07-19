package main

import "fmt"

//冒泡排序
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前", *arr)
	temp := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println("排序后", *arr)
}
func main() {
	arr := [5]int{24, 68, 80, 57, 13}
	BubbleSort(&arr)
}
