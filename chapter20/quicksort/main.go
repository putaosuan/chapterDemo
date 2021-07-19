package main

import (
	"fmt"
	"math/rand"
	"time"
)

//快速排序
func QuickSort(left int, right int, array *[80000]int) {
	l := left
	r := right
	pivot := array[(left+right)/2]
	for l < r {
		for array[l] < pivot {
			l++
		}
		for array[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		array[l], array[r] = array[r], array[l]
		//优化 没看懂
		if array[l] == pivot {
			r--
			//l++
		}
		if array[r] == pivot {
			l++
			//r--
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, array)
	}
	if l < right {
		QuickSort(l, right, array)
	}

}
func main() {
	//arr:=[14]int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	//fmt.Println(arr)
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(90000)
	}
	start := time.Now().Unix()
	QuickSort(0, len(arr)-1, &arr)
	//fmt.Println(arr)
	end := time.Now().Unix()
	fmt.Println(end - start)
}
