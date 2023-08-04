// https://leetcode.com/problems/count-special-quadruplets/description/
package main

// import "sort"

// func countQuadruplets(nums []int) int {
// 	length := len(nums)

// 	sort.Ints(nums)
// 	counter := 0
// 	// left, right := 0, length-1
// 	for right := length - 1; right >= 3; right-- {

// 		// left2, right2 := left+1, right-1
// 		// for left2 < right2 {
// 		// 	sum := nums[left] + nums[left2] + nums[right2]
// 		// 	if sum == nums[right] {
// 		// 		counter += 1
// 		// 		left2 += 1
// 		// 		right2 -= 1
// 		// 	} else if sum < nums[right] {
// 		// 		left2 += 1
// 		// 	}
// 		// }
// 	}
// 	return 0
// }

// func countQuadruplets(nums []int) int {
// 	l := len(nums)
// 	count := 0
// 	sums := map[int]int{}
// 	for left, right := 0, l-1; left < right-2; {
// 		twoSum := nums[left] - nums[right]

// 	}
// 	return 0
// }
