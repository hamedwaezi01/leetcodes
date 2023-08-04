package main

const PRIME = 314159

func hashArray(arr []int, l int) int {
	hash := l
	for i := 0; i < l; i++ {
		hash = hash*PRIME + arr[i]
	}
	return hash
}

func kSum(nums []int, target, k int) [][]int { // The nums array is supposed to be sorted
	length := len(nums)

	result := [][]int{}
	if k == 2 {
		for l, r := 0, length-1; l < r; {
			sum := nums[l] + nums[r]
			if sum == target {
				result = append(result, []int{nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	} else {
		tracker := map[int]bool{}
		for i := 0; i < length; i++ {
			theSums := kSum(nums[i+1:], target-nums[i], k-1)
			for _, arr := range theSums {
				tmp := append(arr, nums[i])
				hash := hashArray(arr, len(arr))
				if _, ok := tracker[hash]; !ok {
					result = append(result, tmp)
					tracker[hash] = true
				}
			}
		}
	}
	return result
}
