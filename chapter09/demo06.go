package main

import (
	"fmt"
	"sort"
)

func main() {
	//map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)

	var keys []int
	for i, _ := range map1 {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _, v := range keys {
		fmt.Printf("map1[%d]=%d\n", v, map1[v])
	}
}
