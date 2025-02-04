/*
ProblemId: 3
ProblemTitle: Longest Substring Without Repeating Characters
ProblemLink: https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given a string s, find the length of the longest substring without repeating characters.

// Brute force here is to find every pair of indices and see how many uniques it has (use a hashmap)
// Another solution is to use a deque and a map
// Keep track of a lo and hi pointer.   Increment hi until the map of chars between low -> hi is unique.
// If there is dup, set low to map[chars[hi]] + 1
func problem3(s string) int {
	if len(s) == 0 {
		return 0
	}
	lastIndex := map[byte]int{s[0]: 0}
	lo := 0
	largest := 0
	for hi := 1; hi < len(s); hi += 1 {
		ch := s[hi]
		if val, ok := lastIndex[ch]; ok && val >= 0 {
			for lo <= lastIndex[ch] {
				lastIndex[s[lo]] = -1
				lo += 1
			}
		}
		lastIndex[ch] = hi
		largest = max(largest, hi-lo)
	}
	return largest + 1
}

func TestProblem3(t *testing.T) {
	assert.Equal(t, problem3(""), 0)
	assert.Equal(t, problem3("abcabcbb"), 3)
	assert.Equal(t, problem3("bbbbb"), 1)
	assert.Equal(t, problem3("pwwkew"), 3)
}
