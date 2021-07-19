package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	//输出看看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	var spareArr []ValNode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	spareArr = append(spareArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				spareArr = append(spareArr, valNode)
			}
		}
	}
	//输出稀疏数组
	fmt.Println("当前的稀疏数组是...")
	for i, v := range spareArr {
		fmt.Printf("%d:%d,%d,%d\n", i, v.row, v.col, v.val)
	}
	//将稀疏数组存盘

	//恢复成原始数组
	//1.打开存盘文件，恢复原始数组

}
