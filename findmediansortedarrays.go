// https://leetcode.com/problems/median-of-two-sorted-arrays/
package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)

	if (l1+l2)%2 == 1 {
		return float64(kthElement(nums1, nums2, 0, l1-1, 0, l2-1, (l1+l2)/2))
	} else {
		return float64(
			kthElement(nums1, nums2, 0, l1-1, 0, l2-1, (l1+l2)/2)+
				kthElement(nums1, nums2, 0, l1-1, 0, l2-1, (l1+l2)/2-1),
		) / 2
	}
}

func kthElement(nums1, nums2 []int, start1, end1, start2, end2, k int) int {
	l1, l2 := end1-start1, end2-start2

	fmt.Println(k, start1, end1, start2, end2, "|", l1, l2, "|", l1/2, l2/2)
	if l1 < 0 {
		return nums2[k-start1]
	}
	if l2 < 0 {
		return nums1[k-start2]
	}
	i1, i2 := start1+(l1/2), start2+(l2/2)
	med1, med2 := nums1[i1], nums2[i2]
	if i1+i2 < k {
		if med1 > med2 {
			return kthElement(nums1, nums2, start1, end1, i2+1, end2, k)
		} else {
			return kthElement(nums1, nums2, i1+1, end1, start2, end2, k)
		}
	} else {
		if med1 > med2 {
			return kthElement(nums1, nums2, start1, i1-1, start2, end2, k)
		} else {
			return kthElement(nums1, nums2, start1, end1, start2, i2-1, k)
		}
	}
}
