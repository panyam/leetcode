package main

import (
	"log"
	"testing"

	"github.com/panyam/leetcode/go/ds"
	"github.com/stretchr/testify/assert"
)

/**
10. Regular Expression Matching - https://leetcode.com/problems/regular-expression-matching/

Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".


Constraints:

1 <= s.length <= 20
1 <= p.length <= 20
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
*/

// BruteForce FIrst:
// if s is empty:
//
//	if p is empty: return True
//	if p.length < 2 or p[1] != '*': return False
//
//	# We *want* this to be // X*.....
//	return match(s, p[2:])
//
// else if p is empty:
//
//	return False
//
// else:
//
//	if p[1] == '*':
//	return match(s[1:], p[1:])

// First the recursive bruteforce option
// Explainable
func isMatch_Recursive(s string, p string) bool {
	var helper func(s string, si, S int, p string, pi, P int) bool
	helper = func(s string, si, S int, p string, pi, P int) bool {
		log.Println("Here on: ", si, pi)
		if si == S { // empty strings
			if pi == P {
				return true // empty pattern
			}
			if pi <= P-2 && p[pi+1] == '*' { // p is at X*
				return helper(s, si, S, p, pi+2, P)
			}
			return false
		} else if pi == P { // empty pattern and non empty string
			return false
		} else if s[si] == p[pi] || p[pi] == '.' { // heads of pat and str match
			if pi <= P-2 && p[pi+1] == '*' {
				return helper(s, si, S, p, pi+2, P) || helper(s, si+1, S, p, pi, P)
			} else {
				return helper(s, si+1, S, p, pi+1, P)
			}
		} else if pi <= P-2 && p[pi+1] == '*' { // s and p dont match so look for wildcard
			return helper(s, si, S, p, pi+2, P)
		}
		return false
	}
	return helper(s, 0, len(s), p, 0, len(p))
}

// Memoized DP -
func isMatch(s string, p string) bool {
	S := len(s)
	P := len(p)
	grid := ds.NewGrid[bool](S+1, P+1)

	// Empty always matches
	grid.Set(S, P, true)

	// Case where pattern is empty - anything but empty string will not match
	for si := S - 1; si >= 0; si-- {
		grid.Set(si, P, false)
	}

	// Case of empty string, pattern only match in some cases
	for pi := P - 1; pi >= 0; pi-- {
		if pi <= P-2 && p[pi+1] == '*' { // p is at X*
			grid.Set(S, pi, grid.Get(S, pi+2))
		} else {
			grid.Set(S, pi, false)
		}
	}

	// rest of the cases
	for si := S - 1; si >= 0; si-- {
		for pi := P - 1; pi >= 0; pi-- {
			res := false
			if s[si] == p[pi] || p[pi] == '.' { // heads of pat and str match
				if pi <= P-2 && p[pi+1] == '*' {
					res = grid.Get(si, pi+2) || grid.Get(si+1, pi)
				} else {
					res = grid.Get(si+1, pi+1)
				}
			} else if pi <= P-2 && p[pi+1] == '*' { // s and p dont match so look for wildcard
				res = grid.Get(si, pi+2)
			}
			grid.Set(si, pi, res)
		}
	}

	// Base case where pattern is empty
	// Any string wont match (unless empty)
	return grid.Get(0, 0)
}

func TestProblem10(t *testing.T) {
	funcs := []func(string, string) bool{isMatch_Recursive, isMatch}

	for idx, fn := range funcs {
		log.Println("Testing with fn: ", idx)
		assert.False(t, fn("aa", "a"))
		assert.True(t, fn("aa", "a*"))
		assert.True(t, fn("ab", ".*"))
	}
}
