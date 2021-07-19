package main

import (
	"fmt"
	"sort"
)

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func twoSum(nums []int, target int) []int {
	var m, n int
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				m = i
				n = j
			}
		}
	}
	var k []int
	k = make([]int, 2)
	k[0] = m
	k[1] = n
	sort.Ints(k)
	return k
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	a := twoSum(nums, target)
	fmt.Println(a)
}
