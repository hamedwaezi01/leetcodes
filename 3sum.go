package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	res := [][]int{}
	set := map[int]bool{}
	for i := 0; i < l; i++ {
		for j, k := i+1, l-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				tmp := []int{nums[i], nums[j], nums[k]}
				k := hashArray(tmp, 3)
				_, ok := set[k]
				if !ok {
					set[k] = true
					res = append(res, tmp)
				}
				// set[k] = tmp
				j += 1
				k -= 1
			} else if sum < 0 {
				j += 1
			} else {
				k -= 1
			}
		}
	}

	return res
}
