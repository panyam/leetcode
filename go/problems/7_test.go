package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem 6 - https://leetcode.com/problems/zigzag-conversion/
// 7. Reverse Integer
//
// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
// Example 1:
//
// Input: x = 123
// Output: 321
//
// Example 2:
//
// Input: x = -123
// Output: -321
//
// Example 3:
//
// Input: x = 120
// Output: 21
func problem7(x int) int {
	if x < 0 {
		return -1 * problem7(-x)
	}
	out := 0
	maxint := (1 << 31) - 1
	for x > 0 {
		rem := x % 10
		x /= 10
		// check if a multiplication would cause this to go > 2^32
		if out > (maxint-rem)/10 {
			return 0
		}
		out = (out * 10) + rem
	}
	return out
}

func TestProblem7(t *testing.T) {
	assert.Equal(t, problem7(123), 321)
	assert.Equal(t, problem7(-123), -321)
	assert.Equal(t, problem7(120), 21)
	assert.Equal(t, problem7(1534236469), 0)
}
