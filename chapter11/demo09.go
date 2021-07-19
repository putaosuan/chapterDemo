package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//声明Student结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}
type students []Student

func (h students) Len() int {
	return len(h)
}

//Less方法就是决定你使用什么标准进行排序
//按成绩从小到大排序
func (h students) Less(i, j int) bool {
	return h[i].Score < h[j].Score
}
func (h students) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func main() {
	var students students
	for i := 1; i <= 10; i++ {
		student := Student{
			Name:  fmt.Sprintf("学生%d", rand.Intn(100)),
			Age:   18,
			Score: float64(rand.Intn(100)),
		}
		students = append(students, student)
	}
	for _, v := range students {
		fmt.Println(v)
	}

	sort.Sort(students)
	fmt.Println("--------排序后----------")
	for _, v := range students {
		fmt.Println(v)
	}
}
