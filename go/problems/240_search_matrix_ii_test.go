package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*
74. Search a 2D Matrix

https://leetcode.com/problems/search-a-2d-matrix/
*/
func searchMatrix2(matrix [][]int, target int) bool {
	R := len(matrix)
	if R == 0 {
		return false
	}
	C := len(matrix[0])

	r, c := 0, C-1
	for c >= 0 && r < R {
		if matrix[r][c] == target {
			return true
		}
		if target < matrix[r][c] {
			c--
		} else {
			r++
		}
	}
	return false
}

func TestProblem240(t *testing.T) {
	assert.True(t, searchMatrix2([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	assert.False(t, searchMatrix2([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}
