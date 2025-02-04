package main

import (
	"testing"

	"github.com/panyam/leetcode/go/algos"
	"github.com/stretchr/testify/assert"
)

/**
34. Find First and Last Position of Element in Sorted Array

https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]
*/

func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	slo, shi := -1, -1
	if hi >= 0 && target >= nums[lo] && target <= nums[hi] {
		// find lowest
		algos.Bisect(lo, hi, func(lo, hi, mid int) (newlo, newhi int, stop bool) {
			if nums[hi] < target || nums[lo] > target {
				return -1, -1, true
			}
			if nums[mid] == target {
				slo = mid
				return lo, mid - 1, false
			} else if nums[mid] < target { // look in right half
				return mid + 1, hi, false
			} else {
				return lo, mid - 1, false
			}
		})

		// find highest (we can avoid this if lo was not found
		algos.Bisect(lo, hi, func(lo, hi, mid int) (newlo, newhi int, stop bool) {
			if nums[hi] < target || nums[lo] > target {
				return -1, -1, true
			}
			if nums[mid] == target {
				shi = mid
				return mid + 1, hi, false
			} else if nums[mid] < target { // look in right half
				return mid + 1, hi, false
			} else {
				return lo, mid - 1, false
			}
		})
	}
	return []int{slo, shi}
}

func TestProblem34(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 0))
	assert.Equal(t, []int{5, 5}, searchRange([]int{1, 2, 3, 4, 5, 6}, 6))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 11))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 3))
}
