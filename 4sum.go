// https://leetcode.com/problems/4sum/
package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	return kSum(nums, target, 4)
}
