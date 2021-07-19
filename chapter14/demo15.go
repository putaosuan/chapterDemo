package main

import (
	"encoding/json"
	"fmt"
)

//演示将json字符串反序列化成struct
type Monster15 struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"nihao_age"'`
	Birthday string
	Sal      float64
	Skill    string
}

func unmarshalStruct() {
	str := "{\"monster_name\":\"牛魔王\",\"nihao_age\":500,\"Birthday\":\"20190923\",\"Sal\":5000,\"Skill\":\"牛魔拳\"}"
	var monser Monster15
	err := json.Unmarshal([]byte(str), &monser)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println("反序列化后，monster=", monser)
}

//演示反序列化Map
func unmarshalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":300,\"name\":\"红孩儿\"}"
	var a map[string]interface{}
	//反序列化Map时，不需要map，因为make被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println("反序列化后，monster=", a)
}

//反序列化Slice切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":7,\"name\":\"jack\"}," +
		"{\"address\":[\"上海\",\"深圳\"],\"age\":7,\"name\":\"mary\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println("反序列化后，monster=", slice)
}
func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
