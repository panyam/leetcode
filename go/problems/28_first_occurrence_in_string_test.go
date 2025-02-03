package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
28.  Find the Index of the First Occurrence in a String - https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string

Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.



Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.


Constraints:

1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.
*/

// BruteForce FIrst:
func strStr(haystack string, needle string) int {
	H := len(haystack)
	N := len(needle)
	for i := 0; i <= H-N; i++ {
		found := true
		for j := 0; j < N && found; j++ {
			if haystack[i+j] != needle[j] {
				found = false
			}
		}

		if found {
			return i
		}
	}
	return -1
}

func TestProblem28(t *testing.T) {
	assert.Equal(t, -1, strStr("leetcode", "leeto"))
	assert.Equal(t, 0, strStr("sadbutsad", "sad"))
	assert.Equal(t, -1, strStr("aaa", "aaaa"))
	assert.Equal(t, 0, strStr("a", "a"))
	assert.Equal(t, 0, strStr("aa", "aa"))

	// Invalid test cases per prob description
	assert.Equal(t, strStr("", "aaaa"), -1)
	assert.Equal(t, strStr("asdfasdf", ""), 0)
}
