package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	l := len(nums)

	sort.Ints(nums)
	fmt.Println(nums)
	val := nums[0] + nums[1] + nums[2]
	diff := getAbs(val - target)

	for i := 0; i < l-2; i++ {
		for left, right := i+1, l-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			abs := getAbs(sum - target)
			// fmt.Println(val, sum, diff, abs)
			if sum == target {
				return sum
			}
			if sum < val {
				left += 1
			} else {
				right -= 1
			}
			if abs < diff {
				diff = abs
				val = sum
			}
		}
	}

	return val
}

func getAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
