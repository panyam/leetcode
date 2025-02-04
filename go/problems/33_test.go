package main

import (
	"testing"

	"github.com/panyam/leetcode/go/algos"
	"github.com/stretchr/testify/assert"
)

/**
33. Search in Rotated Sorted Array

https://leetcode.com/problems/search-in-rotated-sorted-array/

There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
Example 3:

Input: nums = [1], target = 0
Output: -1

*/

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	if hi < 0 {
		return -1
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
			if nums[lo] < nums[mid] { // left half is sorted
				if nums[lo] < target && target < nums[mid] {
					return lo, mid - 1, false
				} else {
					// go to the right half
					return mid + 1, hi, false
				}
			}
			if nums[mid] < nums[hi] {
				if nums[mid] < target && target < nums[hi] {
					return mid + 1, hi, false
				} else {
					return lo, mid - 1, false
				}
			}
		}
		return -1, -1, true
	})
}

func TestProblem33(t *testing.T) {
	assert.Equal(t, 6, search([]int{8, 1, 2, 3, 4, 5, 6, 7}, 6))
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	assert.Equal(t, -1, search([]int{1}, 0))
	assert.Equal(t, 1, search([]int{3, 1}, 1))
	assert.Equal(t, 0, search([]int{3, 1}, 3))
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))
}
