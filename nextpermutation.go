package main

func nextPermutation(nums []int) {
	l, i := len(nums), len(nums)-1
	if l == 1 {
		return
	}

	for i > 0 && nums[i] <= nums[i-1] {
		i -= 1
	}
	i -= 1
	if i < 0 {
		for i = 0; i < l/2; i++ {
			nums[l-1-i], nums[i] = nums[i], nums[l-1-i]
		}
		return
	}
	j := l - 1
	for j > i && nums[i] >= nums[j] {
		j -= 1
	}
	nums[i], nums[j] = nums[j], nums[i]
	i += 1
	for j := 0; j < (l-i)/2; j++ {
		nums[i+j], nums[l-j-1] = nums[l-j-1], nums[i+j]
	}
}
