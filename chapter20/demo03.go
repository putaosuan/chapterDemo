package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

//入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

//显示队列
func (this *CircleQueue) ListQueue() {
	size := this.Size()
	fmt.Println("size=", size)
	if size == 0 {
		fmt.Println("队列为空")
	}
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

//判断队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.head == this.tail
}

//判断队列是否为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}
func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入 add 表示添加数据到队列")
		fmt.Println("1.输入 get 表示从队列获取数据")
		fmt.Println("1.输入 show 表示显示队列")
		fmt.Println("1.输入 exie 表示退出队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入你要添加的数值：")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "show":
			queue.ListQueue()
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数：", val)
			}
		case "exit":
			fmt.Println(queue)
			os.Exit(0)

		}
	}

}
