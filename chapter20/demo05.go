package main

import (
	"fmt"
)

//双向链表
//定义一个HeroNode
type HeroNode2 struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode2
	next     *HeroNode2
}

//
func InsertHeroNode21(head *HeroNode2, newHeroNode *HeroNode2) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp
}
func InsertHeroNode22(head *HeroNode2, newHeroNode *HeroNode2) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}

		temp.next = newHeroNode
	}
}

//显示链表的所有信息
func ListHeroNode2(head *HeroNode2) {
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}
	for {
		fmt.Printf("%d,%s,%s==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
func ListHeroNode23(head *HeroNode2) {

	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	if temp.pre == nil {
		fmt.Println("空空如也...")
		return
	}
	for {
		fmt.Printf("%d,%s,%s==>", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

//删除
func DelHeroNode2(head *HeroNode2, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}

	} else {
		fmt.Println("sorry,要删除的id不存在")
	}
}
func main() {
	head := &HeroNode2{}
	hero1 := &HeroNode2{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode2{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode2{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	ListHeroNode2(head)
	InsertHeroNode22(head, hero3)
	InsertHeroNode22(head, hero1)
	InsertHeroNode22(head, hero2)
	ListHeroNode23(head)
	fmt.Println()
	DelHeroNode2(head, 3)

	ListHeroNode23(head)

}
