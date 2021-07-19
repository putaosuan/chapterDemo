package main

import "fmt"

//转换数组
type www struct {
	a [3][3]int
}

func (w *www) Print() {
	var b [3][3]int

	for i := 0; i < len(w.a); i++ {
		for j := 0; j < len(w.a[i]); j++ {
			b[j][i] = w.a[i][j]
		}
	}
	fmt.Println(b)
}
func main() {
	var w = www{[3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}
	w.Print()

}
