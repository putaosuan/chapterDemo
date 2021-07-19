package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no3"] = "武松"
	a["no4"] = "吴用"
	fmt.Println(a)
}
