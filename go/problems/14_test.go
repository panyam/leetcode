package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem 14: https://leetcode.com/problems/longest-common-prefix/submissions/1513955954/
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	// check cmmon prefixes for both
	for i := range len(first) {
		if first[i] != last[i] {
			if i == 0 {
				return ""
			}
			return first[:i]
		}
	}
	return first
}

func TestProblem14(t *testing.T) {
	assert.Equal(t, longestCommonPrefix([]string{"ab", "a"}), "a")
	assert.Equal(t, longestCommonPrefix([]string{}), "")
	assert.Equal(t, longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl")
	assert.Equal(t, longestCommonPrefix([]string{"dog", "racecar", "car"}), "")
}
