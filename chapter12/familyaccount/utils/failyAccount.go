package utils

import "fmt"

type FamilyAccount struct {
	name    string
	pwd     string
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	details string
	flag    bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		"xiaomi",
		"666888",
		"",
		true,
		10000.0,
		0.0,
		"",
		"收支\t账户余额\t收支金额\t说  明",
		false,
	}
}

//
func (f *FamilyAccount) showDetails() {
	fmt.Println("                   当前收支明细记录                      ")
	if f.flag {
		fmt.Println(f.details)
	} else {
		fmt.Println("当前没有任何收支明细，请添加一笔吧！")
	}
}
func (f *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&f.money)
	f.balance += f.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", f.balance, f.money, f.note)
	f.flag = true
}
func (f *FamilyAccount) pay() {
	fmt.Println("本次支出的金额：")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("金额的余额不足")
		return
	}
	f.balance -= f.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", f.balance, f.money, f.note)
	f.flag = true
}
func (f *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		f.loop = false
	}
}

//转账
func (f *FamilyAccount) zhuanzhang() {
	account := ""
	fmt.Println("请输入对方账户")
	fmt.Scanln(&account)
	fmt.Println("请输入转账金额：")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("金额的余额不足")
		return
	}
	f.balance -= f.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n转账\t%v\t%v\t%v", f.balance, f.money, f.note)
	f.flag = true
}

//主菜单
func (f *FamilyAccount) MainMenu() {
	name := ""
	pwd := ""
	for {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&pwd)
		if name != f.name || pwd != f.pwd {
			fmt.Println("账户或者密码输入错误")
		} else {
			break
		}
	}

	for {
		fmt.Println("\n.................家庭收支记账软件.................")
		fmt.Println("                   1.收支明细                      ")
		fmt.Println("                   2.登记收入                     ")
		fmt.Println("                   3.登记支出                      ")
		fmt.Println("                   4.退   出                      ")
		fmt.Println("                   5.转账                         ")
		fmt.Print("                   请选择（1-5）：")
		fmt.Scanln(&f.key)

		switch f.key {
		case "1":
			f.showDetails()
		case "2":
			f.income()
		case "3":
			f.pay()
		case "4":
			f.exit()
		case "5":
			f.zhuanzhang()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !f.loop {
			break
		}
	}
}
