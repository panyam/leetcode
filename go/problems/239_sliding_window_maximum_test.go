/*
ProblemId: 239
ProblemTitle: Sliding Window Maximum
ProblemLink: https://leetcode.com/problems/sliding-window-maximum/
*/
package main

import (
	"log"
	"testing"

	"github.com/panyam/leetcode/go/ds"
	"github.com/stretchr/testify/assert"
)

// BeginProblemTests

func TestProblem239(t *testing.T) {
	assert.EqualValues(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	assert.EqualValues(t, []int{1}, maxSlidingWindow([]int{1}, 1))
}

// EndProblemTests

// ProblemImpementation:
// Problem impl here fo easy copying

func maxSlidingWindow(nums []int, k int) (out []int) {
	N := len(nums)
	if N == 0 {
		return
	}
	mtq := ds.MonotonicQueue{
		WindowSize: k,
		Less:       func(i, j int) bool { return nums[i] < nums[j] },
	}

	log.Println("Here....")
	for i := range N {
		mtq.PushIndex(i)
		// log.Println("After Pushing: ", i, nums[i], mtq.Indexes(), gfn.Map(mtq.Indexes(), func(index int) int { return nums[index] }))
		if i >= k-1 {
			out = append(out, nums[mtq.MaxIndex()])
		}
	}
	return
}
