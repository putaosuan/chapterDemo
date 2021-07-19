package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//实现对Hero结构体切片的排序
//声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

//声明一个结构体切片
type HeroSlice []Hero

//实现Interface接口
func (h HeroSlice) Len() int {
	return len(h)
}

//Less方法就是决定你使用什么标准进行排序
//按Hero的年龄从小到大排序
func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}
func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

//将Student的切片，按照从大到小的顺序排序
func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	for _, v := range heros {
		fmt.Println(v)
	}

	sort.Sort(heros)
	fmt.Println("--------排序后----------")
	for _, v := range heros {
		fmt.Println(v)
	}
}
