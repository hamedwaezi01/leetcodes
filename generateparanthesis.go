// https://leetcode.com/problems/generate-parentheses/
package main

// This answer is correct but since the order is different from that of test, it throws errors
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	n1 := generateParenthesis(n - 1)
	result := make([]string, len(n1)*3)
	itr := 0
	for _, perm := range n1 {
		result[itr+2] = "(" + perm + ")"
		result[itr+1] = "()" + perm
		result[itr] = perm + "()"
		itr += 3
	}
	return result[1:]
}
