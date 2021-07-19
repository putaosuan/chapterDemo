package main

import (
	"fmt"
)

//定义猫的结构体节点
type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newNode *CatNode) {
	//判断是不是添加一只猫
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head
		return
	}
	temp := head
	for {
		if temp.next == head {
			temp.next = newNode
			newNode.next = head
			break
		}
		temp = temp.next
	}
}
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也")
		return
	}
	for {
		fmt.Println("猫的信息=", temp, "->")
		//fmt.Printf("猫的信息=[id=%d name=%s]->\n",temp.no,temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("空空如也")
		return head
	}
	//只有一个节点
	if temp.next == head {
		if temp.no == id {
			temp.next = nil
		}
		return head
	}
	//将helper定位到链表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//如果有两个以上的节点
	flag := true
	for {
		if temp.next == head { //如果到这，就说明比较到最后一个（最后一个还没有比较）
			break
		}
		if temp.no == id {
			if temp == head { //说明删除的是头节点
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("猫猫%d\n", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫%d\n", id)
		} else {
			fmt.Printf("对不起,没有%d\n", id)
		}
	}
	return head

}
func main() {
	catNode := &CatNode{}
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "jack",
	}
	cat3 := &CatNode{
		no:   3,
		name: "hello",
	}
	InsertCatNode(catNode, cat1)
	InsertCatNode(catNode, cat2)
	InsertCatNode(catNode, cat3)
	ListCircleLink(catNode)
	catNode = DelCatNode(catNode, 1)
	ListCircleLink(catNode)

}
