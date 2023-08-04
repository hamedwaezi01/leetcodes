// https://leetcode.com/problems/trapping-rain-water/
package main

func trap(height []int) int {
	l := len(height)
	if l == 1 {
		return 0
	}
	sum := 0
	notSum := 0
	prev := 0

	for i:=1; i<l; i++ {
		if height[i] < height[prev] {
			notSum += height[i]
		} else {
			tmp := (height[i] * (i-prev-1)) - notSum - ((height[i] - height[prev]) * (i-prev-1))
			sum += tmp
			notSum = 0
			prev = i
		}
	}
	newPrev, prevVal := l, 0
	notSum = 0
	for i:=l-1; i >= prev; i-- {
		if height[i] < prevVal {
			notSum += height[i]
		} else {
			tmp := (height[i] * (newPrev-i-1)) - notSum - ((height[i]-prevVal)* (newPrev-i-1))
			sum += tmp
			notSum = 0
			prevVal = height[i]
			newPrev = i
		}
	}
    return sum
}

