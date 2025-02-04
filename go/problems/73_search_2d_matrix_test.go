package main

import (
	"testing"

	"github.com/panyam/leetcode/go/algos"
	"github.com/stretchr/testify/assert"
)

/*
*
74. Search a 2D Matrix

https://leetcode.com/problems/search-a-2d-matrix/
*/
func searchMatrix(matrix [][]int, target int) bool {
	R := len(matrix)
	if R == 0 {
		return false
	}
	C := len(matrix[0])

	i2rc := func(index int) (row, col int) {
		return int(index / C), index % C
	}

	/*
		rc2i := func(row, col int) int {
			return row*C + col
		}
	*/

	nums := func(index int) int {
		r, c := i2rc(index)
		return matrix[r][c]
	}

	return algos.Bisect(0, R*C-1, func(lo, hi, mid int) (newlo, newhi int, stop bool) {
		mval := nums(mid)
		if mval == target {
			return mid, mid, true
		} else if target < mval {
			return lo, mid - 1, false
		} else {
			return mid + 1, hi, false
		}
	}) >= 0
}

func TestProblem73(t *testing.T) {
	assert.True(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	assert.False(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	assert.False(t, searchMatrix([][]int{{1, 1}}, 2))
	assert.True(t, searchMatrix([][]int{{1, 3}}, 3))
}
