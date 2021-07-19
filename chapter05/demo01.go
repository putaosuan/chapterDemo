package main

import "fmt"

/*
参加百米运动会，如果用时8秒以内进入决赛，否则提示淘汰。
*/
func main() {
	var second float64
	fmt.Println("请输入秒数")
	fmt.Scanln(&second)
	if second < 8 && second > 0 {
		var sex string
		fmt.Println("请输入性别")
		fmt.Scanln(&sex)
		if sex == "男" {
			fmt.Println("进入男子组决赛")
		} else {
			fmt.Println("进入女子组决赛")
		}
	} else {
		fmt.Println("out...")
	}

}
