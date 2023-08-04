// https://leetcode.com/problems/4sum-ii/description/
package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sums := map[int]int{}
	counter := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			key := n3 + n4
			if s, ok := sums[key]; ok {
				sums[key] = s + 1
			} else {
				sums[key] = 1
			}
		}
	}

	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			key := n1 + n2
			if v, ok := sums[-key]; ok {
				counter += v
			}
		}
	}
	return counter
}
