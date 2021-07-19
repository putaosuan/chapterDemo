package main

import "fmt"

func main() {
	/*
		要求：使用一个map来记录monster的信息，name和age，也就是一个monster对应一个map，
		并且妖怪的数量可以动态增加
	*/
	var monsters []map[string]string

	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "孙悟空"
		monsters[1]["age"] = "400"
	}
	newMonsters := make(map[string]string, 2)
	newMonsters["name"] = "火云邪神"
	newMonsters["age"] = "100"
	monsters = append(monsters, newMonsters)
	fmt.Println(monsters)
}
