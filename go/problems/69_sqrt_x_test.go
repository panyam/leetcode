package main

import (
	"testing"

	"github.com/panyam/leetcode/go/algos"
	"github.com/stretchr/testify/assert"
)

/**
69. sqrt(x)  - https://leetcode.com/problems/sqrtx/

Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.


Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
*/

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	return algos.Bisect(0, x, func(lo, hi, mid int) (newlo, newhi int, stop bool) {
		x2 := mid * mid
		if lo == hi || lo == hi-1 || x2 == x {
			return mid, mid, true
		} else if x2 > x {
			return lo, mid, false
		} else {
			return mid, hi, false
		}
	})
}

func TestProblem69(t *testing.T) {
	assert.Equal(t, 1, mySqrt(3))
	assert.Equal(t, 0, mySqrt(0))
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 1, mySqrt(2))
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(8))
}
