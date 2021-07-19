package main

import "fmt"

func main() {
	/*
		某人有100000元，每经过一次路口，需要交费，规则如下：
		当现金大于50000时，每次交费%5
		当现金小于等于50000时，每次交1000
		经过多少次路口？
	*/
	var num = 100000.0
	var count = 0
	for {
		if num > 50000 {
			num *= 0.95
			count++
		} else {
			if num < 1000 {
				break
			}
			num -= 1000
			count++
		}
	}
	fmt.Println("一共经过", count)
}
