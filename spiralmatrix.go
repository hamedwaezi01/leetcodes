// https://leetcode.com/problems/spiral-matrix/
package main

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	total := m * n
	result := make([]int, total)
	l, r, t, b := -1, n, -1, m
	h, v, di, dj := 0, 0, 1, 0
	for k := 0; k < total; k++ {
		result[k] = matrix[v][h]
		h += di
		v += dj
		if di > 0 && h >= r {
			di = 0
			dj = 1
			t++
			h--
			v++
		} else if di < 0 && h <= l {
			di = 0
			dj = -1
			b--
			h++
			v--
		} else if dj < 0 && v <= t {
			di = 1
			dj = 0
			l++
			v++
			h++
		} else if dj > 0 && v >= b {
			di = -1
			dj = 0
			r--
			v--
			h--
		}
	}
	return result
}

// https://leetcode.com/problems/spiral-matrix-ii
func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	i, j, di, dj := 0, 0, 0, 1
	v, h := 0, 0
	total := n * n
	for k := 1; k <= total; k++ {
		matrix[i][j] = k
		i += di
		j += dj
		if dj > 0 && j >= (n-h) {
			dj = 0
			di = 1
			j--
			i++
		} else if dj < 0 && j <= (h-1) {
			dj = 0
			di = -1
			j++
			i--
			h++
		} else if di > 0 && i >= (n-v) {
			dj = -1
			di = 0
			i--
			j--
			v++
		} else if di < 0 && i <= (v-1) {
			dj = 1
			di = 0
			i++
			j++
		}
	}

	return matrix
}

// https://leetcode.com/problems/spiral-matrix-iii/description/
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	total := rows * cols
	matrix := make([][]int, total)
	for i := 0; i < total; i++ {
		matrix[i] = make([]int, 2)
	}
	count := 0 // should reach cols*rows
	r, c, dr, dc := rStart, cStart, 0, 1
	steps, takenSteps := 2, 0
	for count < total {
		// add
		takenSteps++
		if r < rows && r >= 0 && c < cols && c >= 0 {
			matrix[count][0] = r
			matrix[count][1] = c
			count++
		}
		// next
		r += dr
		c += dc
		if takenSteps < steps {
			continue
		}
		takenSteps = 1

		if dc != 0 {
			c -= dc
			r += dc
			dr = dc
			dc = 0
		} else if dr != 0 {
			steps++
			r -= dr
			c -= dr
			dc = -dr
			dr = 0
		}
	}

	return matrix
}

/** https://leetcode.com/problems/spiral-matrix-iv/description/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	total := n * m
	count := 0
	matrix := make([][]int, total)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	r, c, dr, dc := 0, 0, 0, 1
	v, h := 0, 0
	for count <= total {
		if head == nil {
			matrix[r][c] = -1
		} else {
			matrix[r][c] = head.Val
		}
		head = head.Next
		r += dr
		c += dc

		if dc > 0 && c >= (n-h) {
			dc = 0
			dr = 1
			c--
			r++
		} else if dc < 0 && c <= (h-1) {
			dc = 0
			dr = -1
			c++
			r--
		} else if dr > 0 && r >= (m-v) {
			dc = -1
			dr = 0
			c--
			r--
		} else if dr < 0 && r <= (v-1) {
			dc = 1
			dr = 0
			c++
			r++
		}
	}

	return matrix
}
