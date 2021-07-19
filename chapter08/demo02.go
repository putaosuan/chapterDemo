package main

import "fmt"

func main() {
	//有一个数列：白眉鹰王，金毛狮王，紫衫龙王，青翼蝠王
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼辐网"}
	var heroname = ""
	fmt.Println("请输入要查找的人名")
	fmt.Scanln(&heroname)
	//顺序查找
	for i := 0; i < len(names); i++ {
		if heroname == names[i] {
			fmt.Println("找到了")
			break
		} else if i == len(names)-1 {
			fmt.Println("没有找到")
		}
	}
	//第二种
	index := -1
	for i := 0; i < len(names); i++ {
		if heroname == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Println("zhaodoal下标：", index)
	} else {
		fmt.Println("meizhaodao")
	}
}
