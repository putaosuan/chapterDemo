package main

import (
	"errors"
	"fmt"
)

//使用数组来模拟一个栈
type Stack struct {
	MaxTop int
	Top    int
	arr    [5]int //数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

//遍历栈
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("栈空")
		return
	}
	fmt.Println("栈的情况如下")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("栈空")
		return 0, errors.New("stack empty")
	}
	//先取值
	val = this.arr[this.Top]
	this.Top--
	return val, nil

}
func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.List()
	val, _ := stack.Pop()
	fmt.Println(val)
	val, _ = stack.Pop()
	fmt.Println(val)
	stack.List()

}
