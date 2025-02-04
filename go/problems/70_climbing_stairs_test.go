/*
ProblemId: 70
ProblemTitle: Climibing Stairs
ProblemLink: https://leetcode.com/problems/climbing-stairs/
*/
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BeginProblemTests

func TestProblem70(t *testing.T) {
	runTest := func(n, expected int) {
		assert.Equal(t, expected, climbStairsIterative(n))
		assert.Equal(t, expected, climbStairsWithCache(n, nil))
	}
	runTest(45, 1836311903)
	runTest(10, 89)
	runTest(2, 2)
	runTest(1, 1)
	runTest(0, 0)
}

// EndProblemTests

// ProblemImpementation:
// Problem impl here fo easy copying
func climbStairs(n int) int {
	return climbStairsIterative(n)
}

// Quick iterative version
func climbStairsIterative(n int) int {
	cache := make(map[int]int)
	for i := range n + 1 {
		if i <= 2 {
			cache[i] = i
		} else {
			cache[i] = cache[i-1] + cache[i-2]
		}
	}
	return cache[n]
}

// Memoized DB
func climbStairsWithCache(n int, cache map[int]int) int {
	if cache == nil {
		cache = make(map[int]int)
		for i := range n + 1 {
			cache[i] = -1
			if i <= 2 {
				cache[i] = i
			}
		}

	}
	if out, ok := cache[n]; !ok || out < 0 {
		cache[n] = climbStairsWithCache(n-1, cache) + climbStairsWithCache(n-2, cache)
	}
	return cache[n]
}
