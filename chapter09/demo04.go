package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "深圳"
	fmt.Println(cities)
	//修改
	cities["no3"] = "山东"
	fmt.Println(cities)
	//删除 delete,是一个内置函数，如果key存在，就删除key-value，如果不存在，不操作但不不会报错
	delete(cities, "no1")
	delete(cities, "no4")
	fmt.Println(cities)
	//如果我们希望一次性删除所有的key
	//1.遍历所有的key，一次性删除
	//2.直接make一个新的空间,
	//cities=make(map[string]string)
	//fmt.Println(cities)

	//map 查找
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no2,key值为%v\n", val)
	} else {
		fmt.Println("不存在")
	}
}
