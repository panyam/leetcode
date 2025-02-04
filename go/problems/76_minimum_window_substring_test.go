/*
ProblemId: 76
ProblemTitle: Minimum Window Substring
ProblemLink: https://leetcode.com/problems/minimum-window-substring/
*/
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BeginProblemTests

func TestProblem76(t *testing.T) {
	assert.Equal(t, "abcabc", minWindow("abcabc", "cacabb"))
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "a", minWindow("a", "a"))
	assert.Equal(t, "", minWindow("a", "aa"))
	assert.Equal(t, "aa", minWindow("aa", "aa"))
}

// EndProblemTests

// ProblemImpementation:
// Problem impl here fo easy copying
func minWindow(s, t string) string {
	has := make(map[byte]bool)
	counts := make(map[byte]int)
	currcnt := make(map[byte]int)
	for _, ch := range []byte(t) {
		has[ch] = true
		counts[ch] += 1
		currcnt[ch] = 0
	}
	// log.Println("Has, Counts: ", has, counts)

	// Where the current "match" starts
	// will be set when first char in T is found and updated subsequently
	matchStart, matchEnd := -1, -1
	bytes := []byte(s)
	S := len(bytes)
	deque := make([]int, S)
	// low and hi pointers into the dequeue
	dl, dh := 0, 0

	// How many chars are to be matched
	T := len(t)
	rem := T

	for hi := 0; hi < S; {
		ch := bytes[hi]
		if !has[ch] {
			hi++
			continue
		}

		deque[dh] = hi
		dh += 1

		// add to the sliding window and increment right ptr

		// increment usage of current char
		currcnt[ch] += 1
		if currcnt[ch] == counts[ch] {
			// a character quota has filled up
			// any more characters added should not increment this
			rem -= counts[ch]
		}

		// found a full match
		// log.Println("ch, dl, dh, deque, hi, rem: ", ch, dl, dh, deque, hi, rem)
		for rem == 0 && dl < dh {
			// we have a match, so record it
			if matchStart == -1 || deque[dh-1]-deque[dl] < matchEnd-matchStart {
				matchStart, matchEnd = deque[dl], deque[dh-1]
			}

			// now move the left ptr forward and pop it
			leftch := bytes[deque[dl]]
			currcnt[leftch] -= 1
			// log.Println("Leftch, currcnt, counts: ", leftch, currcnt[leftch], counts[leftch])
			dl += 1
			if currcnt[leftch] == counts[leftch]-1 {
				rem += counts[leftch]
			}
		}
		// log.Println("MS, ME: ", matchStart, matchEnd)
		hi += 1
	}
	if matchStart < 0 {
		return ""
	}
	return s[matchStart : matchEnd+1]
}
