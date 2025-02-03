package main

import (
	"testing"

	"github.com/panyam/leetcode/go/algos"
	"github.com/stretchr/testify/assert"
)

/**
35. Search Insert Position - https://leetcode.com/problems/search-insert-position/

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4


Constraints:

1 <= nums.length <= 10^4
-104 <= nums[i] <= 10^4
nums contains distinct values sorted in ascending order.
-104 <= target <= 10^4
*/

func searchInsert(nums []int, target int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	if target < nums[0] {
		return 0
	}
	if target > nums[N-1] {
		return N
	}
	return algos.Bisect(0, N-1, func(lo, hi, mid int) (newlo, newhi int, stop bool) {
		// log.Println("Target, lo, mid, hi: ", target, lo, mid, hi)
		if nums[mid] == target {
			return mid, mid, true
		} else if lo == hi {
			if nums[lo] < target {
				return lo + 1, lo + 1, true
			} else {
				return lo, lo, true
			}
		} else if nums[mid] < target { // right half
			return mid + 1, hi, false
		} else { // left half
			return lo, mid, false
		}
	})
}

func TestProblem35(t *testing.T) {
	assert.Equal(t, 0, searchInsert([]int{1, 3}, 0))
	assert.Equal(t, 0, searchInsert([]int{1, 3, 5, 6}, 0))
	assert.Equal(t, 0, searchInsert([]int{1, 3, 5}, 1))
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
	assert.Equal(t, 3, searchInsert([]int{3, 5, 7, 9, 10}, 8))
}
