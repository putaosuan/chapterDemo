package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

//从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	return this.array[this.front], err
}

//显示队列
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}
func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "show":
			queue.ShowQueue()
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数：", val)
			}
		case "exit":
			os.Exit(0)

		}
	}
}
