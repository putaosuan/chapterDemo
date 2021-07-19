package main

import (
	"fmt"
	"os"
)

//定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) ShowMe() {
	fmt.Printf("链表%d找到该雇员%d\n", this.Id%7, this.Id)
}

type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head
	var pre *Emp = nil
	if cur == nil {
		this.Head = emp
		return
	}
	for {
		if cur != nil {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	if pre == nil {
		this.Head = emp
		emp.Next = cur
		return
	}
	pre.Next = emp
	emp.Next = cur
}

//显示
func (this *EmpLink) ShowAll(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d为空\n", no)
		return
	}
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员=%d 名字=%s->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}
func (this *EmpLink) FindByid(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) Insert(emp *Emp) {
	linkNo := this.HashFun(emp.Id)
	this.LinkArr[linkNo].Insert(emp)
}
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowAll(i)
	}
}
func (this *HashTable) HashFun(id int) int {
	return id % 7
}
func (this *HashTable) FindByid(id int) *Emp {
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindByid(id)
}
func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("============雇员系统菜单=============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入Id号")
			fmt.Scanln(&id)
			emp := hashtable.FindByid(id)
			if emp == nil {
				fmt.Println("不存在")
			} else {
				emp.ShowMe()
			}

		case "exit":
			os.Exit(0)
		}

	}
}
