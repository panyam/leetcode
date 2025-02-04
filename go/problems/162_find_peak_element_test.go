package main

import (
	"testing"

	"github.com/panyam/leetcode/go/algos"
	"github.com/stretchr/testify/assert"
)

/*
*
162. Find Peak Element

https://leetcode.com/problems/find-peak-element/
*/
func findPeakElement(nums []int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 0
	}

	peakDir := func(i int) int {
		if i == 0 {
			if nums[i] > nums[i+1] {
				return 0
			} else {
				return 1
			}
		}
		if i == N-1 {
			if nums[i] > nums[i-1] {
				return 0
			} else {
				return -1
			}
		}
		if nums[i-1] < nums[i] && nums[i+1] < nums[i] {
			return 0
		}

		// now other cases
		if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			return 1
		}
		return -1
	}

	return algos.Bisect(0, N-1, func(lo, hi, mid int) (newlo, newhi int, stop bool) {
		dir := peakDir(mid)
		if dir == 0 {
			return mid, mid, true
		} else if dir == 1 { // peak is on the right as it is sloping up to the right
			return mid + 1, hi, false
		}
		return lo, mid - 1, false
	})
}

func TestProblem162(t *testing.T) {
	assert.Equal(t, 2, findPeakElement([]int{1, 2, 3, 1}))
	assert.Equal(t, 5, findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
