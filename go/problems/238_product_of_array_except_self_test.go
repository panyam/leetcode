package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Problem 238 - Product of Array Except Self - https://leetcode.com/problems/product-of-array-except-self/description/

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.

Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/
func productExceptSelf(nums []int) []int {
	N := len(nums)
	out := []int{0}
	accum := nums[0]
	for i := range N - 1 {
		out = append(out, accum)
		accum *= nums[i+1]
	}

	accum = nums[N-1]
	for i := N - 2; i > 0; i-- {
		out[i] *= accum
		accum *= nums[i]
	}
	out[0] = accum
	return out
}

func TestProblem238(t *testing.T) {
	assert.Equal(t, productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6})
	assert.Equal(t, productExceptSelf([]int{-1, 1, 0, -3, 3}), []int{0, 0, 9, 0, 0})
}
