package main

import "fmt"

//小孩的结构体
type Boy struct {
	No   int
	Next *Boy
}

//
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("num值不对")
		return first
	}
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first

}

//线是单向的环形列表
func ShowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("空空如也")
		return
	}
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d->", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next

	}
	fmt.Println()
}

//
func PlayGame(first *Boy, startNo int, CountNum int) {
	if first.Next == nil {
		fmt.Println("空的链表")
		return
	}
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	for {
		for i := 1; i <= CountNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号%d出圈\n", first.No)
		first = first.Next
		tail.Next = first
		if tail == first {
			break
		}
	}
	fmt.Printf("最后小孩编号%d出圈\n", first.No)
}
func main() {
	first := AddBoy(50)
	ShowBoy(first)
	PlayGame(first, 20, 31)

}
