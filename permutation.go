// https://leetcode.com/problems/permutations/description/ TODO
package main

import "fmt"

// func permute(nums []int) [][]int {
// 	l := len(nums)
// 	n := factorial(l)
// 	result := make([][]int, n)
// 	for i := range result {
// 		result[i] = make([]int, l)
// 	}

// 	// for i := 0; i < n; i++ {
// 	// 	temp := make([]int, l)
// 	// 	for j := 0; j < l; j++ {
// 	// 		temp[j] = nums[((i/l)+j)%l]
// 	// 	}
// 	// 	fmt.Println(temp)
// 	// 	result[i] = temp
// 	// 	//  = nums[i%l]
// 	// }
// 	return result
// }

func permute(nums []int) [][]int {
	l := len(nums)
	// if l == 2 {
	// 	return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	// }
	if l == 1 {
		return [][]int{nums}
	}
	fmt.Println()
	perms := permute(nums[1:])
	// n1 := factorial(l - 1)
	result := make([][]int, factorial(l))
	for i := range result {
		result[i] = make([]int, l)
	}
	for i, perm := range perms {
		for j, v := range perm {
			fmt.Println(i, j, " | ", i*(l-1)+(j%l), v)
			// result[i*(l-1)+(j%l)][] = v
		}
	}
	// for i := 0; i < l; i++ {
	// 	temp := make([]int, l)
	// 	temp[i] = nums[0]
	// 	j := i + 1
	// 	for j != i {
	// 		temp[]
	// 		j = (j + 1) % l
	// 	}
	// 	// for j := 0; j < l; j++ {
	// 	// 	temp[j] = perms[i/n1][]
	// 	// }
	// }
	return result
}
