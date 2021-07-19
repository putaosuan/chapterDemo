package main

import (
	"chapterDemo/chapter13/customerManage/model"
	"chapterDemo/chapter13/customerManage/service"
	"fmt"
)

type custerView struct {
	key           string
	loop          bool
	custerService *service.CustomerService
}

//显示所有客户信息
func (c *custerView) list() {
	customers := c.custerService.List()
	fmt.Println("----------------客户列表----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("----------------客户列表完成----------------\n\n")
}

//添加客户信息
func (c *custerView) add() {
	fmt.Println("----------------添加客户----------------")
	fmt.Print("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱")
	email := ""
	fmt.Scanln(&email)
	cusmoter := model.NewCustomer2(name, gender, age, phone, email)
	if c.custerService.Add(cusmoter) {
		fmt.Println("添加用户成功")
	} else {
		fmt.Println("添加用户失败")
	}
}

//删除用户输入的ID
func (c *custerView) delete() {
	fmt.Println("----------------删除客户----------------")
	fmt.Println("请选择删除客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除y/n：")
	choice := ""
	for {
		fmt.Scan(&choice)
		if choice == "y" || choice == "Y" {
			if c.custerService.Delete(id) {
				fmt.Println("------------删除完成----------")
				return
			} else {
				fmt.Println("删除失败，输入的ID号不存在")
			}
		} else if choice == "n" || choice == "N" {
			return
		} else {
			fmt.Println("输入不正确，请重新输入y/n")
		}
	}

}

//退出
func (c *custerView) exit() {
	fmt.Println("确认是否退出y/n")
	for {
		fmt.Scanln(&c.key)
		if c.key == "y" || c.key == "n" {
			break
		} else {
			fmt.Println("你的输入有误,请重新输入")
		}
	}
	if c.key == "y" {
		c.loop = false
	}

}

//修改信息
func (c *custerView) upDate() {
	fmt.Println("------------修改客户----------")
	fmt.Println("请选择修改客户编号(-1退出)")
	id := 0
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := c.custerService.FindById(id)
	if index != -1 {
		fmt.Printf("姓名（%v）", c.custerService.List()[index].Name)
		name := ""
		fmt.Scanln(&name)
		if name == "" {
			name = c.custerService.List()[index].Name
		}
		fmt.Printf("性别（%v）:", c.custerService.List()[index].Gender)
		gender := ""
		fmt.Scanln(&gender)
		if gender == "" {
			gender = c.custerService.List()[index].Gender
		}
		fmt.Printf("年龄（%v）:", c.custerService.List()[index].Age)
		age := 0
		fmt.Scanln(&age)
		if age == 0 {
			age = c.custerService.List()[index].Age
		}
		fmt.Printf("电话（%v）:", c.custerService.List()[index].Phone)
		phone := ""
		fmt.Scanln(&phone)
		if phone == "" {
			phone = c.custerService.List()[index].Phone
		}
		fmt.Printf("邮箱（%v）:", c.custerService.List()[index].Email)
		email := ""
		fmt.Scanln(&email)
		if email == "" {
			email = c.custerService.List()[index].Email
		}
		c.custerService.UpDate(index, name, gender, age, phone, email)
		fmt.Println("------------修改完成----------")
	} else {
		fmt.Println("客户编号不存在")
	}

}
func (c *custerView) mainMenu() {
	for {
		fmt.Println("----------------客户信息管理软件----------------")
		fmt.Println("                 1.添加客户")
		fmt.Println("                 2.修改客户")
		fmt.Println("                 3.删除客户")
		fmt.Println("                 4.客户列表")
		fmt.Println("                 5.退   出")
		fmt.Println("请选择（1-5）：")
		fmt.Scanln(&c.key)

		switch c.key {
		case "1":
			c.add()
		case "2":
			c.upDate()
		case "3":
			c.delete()
		case "4":
			c.list()
		case "5":
			c.exit()
		default:
			fmt.Println("你的输入有误，请重新输入。。。")
		}
		if !c.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理软件")
}
func main() {
	custerView := custerView{
		key:  "",
		loop: true,
	}
	custerView.custerService = service.NewCustomerService()
	custerView.mainMenu()
}
