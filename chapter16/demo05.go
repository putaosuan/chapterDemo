package main

import "fmt"

func main() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "天津"
	m2 := make(map[string]string, 10)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"
	mapChan <- m1
	mapChan <- m2
	fmt.Println(mapChan)
}
