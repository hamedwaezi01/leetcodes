// https://leetcode.com/problems/jump-game-iii/description/
package main

// func canReach(arr []int, start int) bool {
// 	if 0 <= start && start < len(arr) && arr[start] >= 0 {
// 		arr[start] = -arr[start]
// 		return arr[start] == 0 || canReach(arr, start+arr[start]) || canReach(arr, start-arr[start])
// 	}
// 	return false
// }

func canReach(arr []int, start int) bool {
	l := len(arr)
	stack := make([]int, l)
	stack[0] = start
	stackPin := 0
	for stackPin >= 0 {
		index := stack[stackPin]
		stackPin -= 1
		if arr[index] == 0 {
			return true
		}

		if arr[index] < 0 {
			continue
		}

		if index+arr[index] < l {
			stackPin += 1
			stack[stackPin] = arr[index] + index
		}
		if index-arr[index] >= 0 {
			stackPin += 1
			stack[stackPin] = index - arr[index]
		}
		arr[index] *= -1
	}

	return false
}
