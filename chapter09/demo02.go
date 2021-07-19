package main

import "fmt"

func main() {
	//第一种使用方式
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no3"] = "武松"
	a["no4"] = "吴用"
	//第二种
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no1"] = "天津"
	cities["no1"] = "上海"
	//第三种
	hreos := map[string]string{
		"hreo1": "宋江",
		"hreo2": "吴用",
		"hreo3": "卢俊义",
	}
	fmt.Println(hreos)
}
