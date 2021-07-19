package main

import "fmt"

func SetWay(mymap *[8][7]int, i int, j int) bool {
	if mymap[6][5] == 2 {
		return true
	} else {
		if mymap[i][j] == 0 {
			mymap[i][j] = 2
			if SetWay(mymap, i+1, j) {
				return true
			} else if SetWay(mymap, i, j+1) {
				return true
			} else if SetWay(mymap, i-1, j) {
				return true
			} else if SetWay(mymap, i, j-1) {
				return true
			} else {
				mymap[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}
func main() {
	var myMap [8][7]int
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	SetWay(&myMap, 1, 1)
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf("%d\t", myMap[i][j])
		}
		fmt.Println()
	}
}
