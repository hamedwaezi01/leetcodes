package main

import (
	"fmt"
)

func main() {
	// driver3SumClosest()
	// driver3Sum()
	// driver4Sum()
	// driverPhoneNumber()
	// driver4SumCount()
	// driverGenerateParanthesis()
	// driverNextPermutation()
	// driverPermute() //???
	// driverProductExceptSelf()
	// driverFindMedianSortedArrays()
	// driverRemoveNthFromEnd()
	driverjumpgame()
}

func driverjumpgame() {
	// fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	// fmt.Println(jump([]int{2, 3, 0, 1, 4}))

	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	// fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))
}

func driverRemoveNthFromEnd() {
	input := []int{1, 2, 3, 4}
	head := &ListNode{Val: input[0]}
	current := head
	for _, v := range input[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	current = head
	removeNthFromEnd(head, len(input))
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
}

func driverFindMedianSortedArrays() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	// fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	// fmt.Println(findMedianSortedArrays([]int{0, 0, 0, 0, 0}, []int{-1, 0, 0, 0, 0, 0, 1}))
}

func driverProductExceptSelf() {
	// fmt.Println(productExceptSelf([]int{-1,1,0,-3,3}))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func driverWaterTrap() {
	// fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap([]int{4, 2, 3}))
}

func driverNextPermutation() {
	nextPermutation([]int{1, 5, 1})

	// nextPermutation([]int{1, 3, 2})
	// nextPermutation([]int{4, 3, 2, 1})
}

func driverPermute() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func driverSwapNodesInPair() {
	input := []int{1, 2, 3, 4}
	head := ListNode{Val: input[0]}
	current := head
	for _, v := range input[1:] {
		current.Next = &ListNode{Val: v}
		current = *current.Next
	}

	fmt.Println(swapPairs(&head))
}

func driverGenerateParanthesis() {
	fmt.Println(generateParenthesis(2))
}

func driverPhoneNumber() {
	fmt.Println(letterCombinations("234"))
}

func driver3SumClosest() {
	// fmt.Println((1<<32 - 1))
	fmt.Println(threeSumClosest([]int{833, 736, 953, -584, -448, 207, 128, -445, 126, 248, 871, 860, 333, -899, 463, 488, -50, -331, 903, 575, 265, 162, -733, 648, 678, 549, 579, -172, -897, 562, -503, -508, 858, 259, -347, -162, -505, -694, 300, -40, -147, 383, -221, -28, -699, 36, -229, 960, 317, -585, 879, 406, 2, 409, -393, -934, 67, 71, -312, 787, 161, 514, 865, 60, 555, 843, -725, -966, -352, 862, 821, 803, -835, -635, 476, -704, -78, 393, 212, 767, -833, 543, 923, -993, 274, -839, 389, 447, 741, 999, -87, 599, -349, -515, -553, -14, -421, -294, -204, -713, 497, 168, 337, -345, -948, 145, 625, 901, 34, -306, -546, -536, 332, -467, -729, 229, -170, -915, 407, 450, 159, -385, 163, -420, 58, 869, 308, -494, 367, -33, 205, -823, -869, 478, -238, -375, 352, 113, -741, -970, -990, 802, -173, -977, 464, -801, -408, -77, 694, -58, -796, -599, -918, 643, -651, -555, 864, -274, 534, 211, -910, 815, -102, 24, -461, -146},
		-7111))
	// fmt.Println(threeSumClosest([]int{-100, 0, 20, -5, -8}, -10))
	// fmt.Println(threeSumClosest([]int{1, 3, 4, 7, 8, 9}, 15))
}

func driver4SumCount() {
	// fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	fmt.Println(fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}))
}
func driver4Sum() {
	// fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
	// fmt.Println(append([]int{1, 2, 3}, []int{4, 5, 6}...))
}

func driver3Sum() {
	// fmt.Println(hashArray([]int{-1, 5, 2}, 3))
	// fmt.Println(threeSum([]int{3, 0, -2, -1, 1, 2}))
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
