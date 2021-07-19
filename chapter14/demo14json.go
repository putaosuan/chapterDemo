package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"nihao_age"'`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	var monster = Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "20190923",
		Sal:      5000.0,
		Skill:    "牛魔拳",
	}
	date, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败,err=", err)
	}
	fmt.Println(string(date))
}
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 300
	a["address"] = "洪崖洞"
	date, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败,err=", err)
	}
	fmt.Println(string(date))
}
func testSlice() {
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "北京"
	slice = append(slice, m1)
	m2 := make(map[string]interface{})
	m2["name"] = "mary"
	m2["age"] = 7
	m2["address"] = [2]string{"上海", "深圳"}
	slice = append(slice, m2)

	date, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败,err=", err)
	}
	fmt.Println(string(date))
}
func main() {
	//演示将结构体，map，切片进行序列化
	testStruct()
	testMap()
	testSlice()

}
