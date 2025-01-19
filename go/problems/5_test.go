package main

import (
	"testing"

	"github.com/panyam/leetcode/go/ds"
	"github.com/stretchr/testify/assert"
)

// Problem 5 - https://leetcode.com/problems/longest-palindromic-substring/
// Given a string s, find the length of the longest palindromic subsequence in s.
//
// Ideas
// Brute force:
//   - For every pair of ranges - see if it is a palindrome and if it is a largest
//   - Complexity - O(n^2) iterations and palindrome check is O(n). Overall O(n^3)
//
// Observation:
//
// For i,j where j - 1 >= 2, s[i..j] is a palindrome if s[i + 1 ... j - 1] is a palindrome and s[i] == s[j]
//
// Memozing this we could bring this to O(n^2) (and O(n^2) space)

func problem5_bruteforce(s string) string {
	N := len(s)
	isPalindrome := func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i += 1
			j -= 1
		}
		return true
	}

	if N <= 1 {
		return s
	}
	largest := 0
	li, lj := -1, -1
	for j := N - 1; j >= 0; j-- {
		for i := 0; i <= j; i += 1 {
			if 1+j-i > largest && isPalindrome(i, j) {
				largest = 1 + j - i
				li, lj = i, j
			}
		}
	}
	return s[li : lj+1]
}

func isPalindrome(s string, i, j int, cache *ds.Grid[int]) bool {
	if i == j || (i == j-1 && s[i] == s[j]) {
		return true
	}

	res := cache.SafeGet(i, j, 0)
	if res == 0 {
		if s[i] != s[j] {
			res = -1
		} else {
			if isPalindrome(s, i+1, j-1, cache) {
				res = 1
			} else {
				res = -1
			}
		}
		cache.SafeSet(i, j, res)
	}
	return res > 0
}

func problem5_dp(s string) string {
	N := len(s)

	cache := ds.NewGrid[int](N, N)
	if N <= 1 {
		return s
	}
	largest := 0
	li, lj := -1, -1
	for j := N - 1; j >= 0; j-- {
		for i := 0; i <= j; i += 1 {
			if 1+j-i > largest && isPalindrome(s, i, j, cache) {
				largest = 1 + j - i
				li, lj = i, j
			}
		}
	}
	return s[li : lj+1]
}

func TestProblem5(t *testing.T) {
	assert.Equal(t, problem5_dp(""), "")
	assert.Equal(t, problem5_dp("ac"), "c")
	assert.Equal(t, problem5_dp("cbbd"), "bb")
	assert.Equal(t, problem5_dp("babad"), "aba")

	assert.Equal(t, problem5_bruteforce(""), "")
	assert.Equal(t, problem5_bruteforce("ac"), "c")
	assert.Equal(t, problem5_bruteforce("cbbd"), "bb")
	assert.Equal(t, problem5_bruteforce("babad"), "aba")
}
