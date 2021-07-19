package main

import "fmt"

type Box struct {
	len    float64
	width  float64
	height float64
}

//声明一个方法获取立方体的体积
func (b *Box) getVolumn() float64 {
	return b.len * b.width * b.height
}
func main() {
	var b Box
	fmt.Println("请输入长度：")
	fmt.Scanln(&b.len)
	fmt.Println("请输入宽度：")
	fmt.Scanln(&b.width)
	fmt.Println("请输入高度：")
	fmt.Scanln(&b.height)
	fmt.Println(b.getVolumn())
}
