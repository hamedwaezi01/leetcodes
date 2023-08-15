// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	last := 0
	current := head
	for current != nil {
		last += 1
		current = current.Next
	}
	if last == n {
		return head.Next
	}

	current = head
	for i := 0; i < last-n-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	return head
}
