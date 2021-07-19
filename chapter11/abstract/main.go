package main

import "fmt"

//定义一个结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//存款
func (acc *Account) Deposite(money float64, pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	acc.Balance += money
	fmt.Println("存款成功")
}

//取款
func (acc *Account) WithDraw(money float64, pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money <= 0 || money > acc.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	acc.Balance -= money
	fmt.Println("取款成功")
}

//查询余额
func (acc *Account) Query(pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号=%v,余额=%v\n", acc.AccountNo, acc.Balance)
}
func main() {
	account := Account{"gs1111", "666666", 100.0}
	//account.Query("666666")
	//account.Deposite(200,"666666")
	//account.Query("666666")
	//account.WithDraw(50,"6666666")
	//account.Query("666666")
	var pwd string
	fmt.Println("请输入密码")
	fmt.Scanln(&pwd)
	if pwd != account.Pwd {
		fmt.Println("密码输入错误")
		return
	}
	n := 0
	var count float64
	for {
		fmt.Println("请输入命令选项 1：查询 2：存款 3：取款 4：结束查询")
		fmt.Scanln(&n)
		switch n {
		case 1:
			account.Query(pwd)
		case 2:
			fmt.Println("请输入存款金额")
			fmt.Scanln(&count)
			account.Deposite(count, pwd)
			account.Query(pwd)
		case 3:
			fmt.Println("请输入取款金额")
			fmt.Scanln(&count)
			account.WithDraw(count, pwd)
			account.Query(pwd)
		case 4:
			return
		default:
			fmt.Println("输入模式有误")
		}
	}

}
