// https://leetcode.com/problems/product-of-array-except-self/
package main

import "fmt"

func productExceptSelf(nums []int) []int {
    l := len(nums)
	product := 1
	zeros := 0

	for i:=0; i<l; i++ {
		if nums[i] == 0 {
			zeros += 1
		}else {
			product *= nums[i]
		}
	}


	for i:=0; i<l; i++ {
		if nums[i] == 0 {
			if zeros > 1 {
				nums[i] = 0
			} else {
				nums[i] = product
			}
		} else {
			if zeros > 0 {
				nums[i] = 0
			} else {
				nums[i] = product / nums[i]
			}
		}
	}

    return nums
}

/*
		non-zero	zero
shit		0		 *

no-shit	 	*		NaN
*/