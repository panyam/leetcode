package main

import (
	"testing"

	"github.com/panyam/leetcode/go/algos"
	"github.com/stretchr/testify/assert"
)

/**
81. Search in Rotated Sorted Array II (Without Distinct)

There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

You must decrease the overall operation steps as much as possible.



Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
*/

func search81(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	if hi < 0 {
		return false
	}
	// find lowest
	return algos.Bisect(lo, hi, func(lo, hi, mid int) (newlo, newhi int, stop bool) {
		// log.Println("Here, target, l, m, h: ", target, lo, mid, hi)
		if nums[mid] == target {
			return mid, mid, true
		} else if nums[lo] == target {
			return lo, lo, true
		} else if nums[hi] == target {
			return hi, hi, true
		} else {
			if nums[lo] == nums[mid] && nums[mid] == nums[hi] {
				return lo + 1, hi - 1, false
			}
			if nums[lo] <= nums[mid] { // left half is sorted
				if nums[lo] < target && target < nums[mid] {
					// and it contains it
					return lo, mid - 1, false
				} else {
					// go to the right half
					return mid + 1, hi, false
				}
			}
			if nums[mid] <= nums[hi] { // right half is rotated
				if nums[mid] < target && target < nums[hi] {
					// and it contains it
					return mid + 1, hi, false
				} else {
					return lo, mid - 1, false
				}
			}
		}
		return -1, -1, true
	}) >= 0
}

func TestProblem81(t *testing.T) {
	assert.True(t, search81([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	assert.False(t, search81([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	assert.True(t, search81([]int{1, 0, 1, 1, 1}, 0))
	assert.True(t, search81([]int{2, 2, 2, 0, 0, 1}, 0))
}
