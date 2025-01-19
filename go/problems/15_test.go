package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem 15 -
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
func threeSum(nums []int) (out [][]int) {
	N := len(nums)
	sort.Ints(nums)

	for i := 0; i < N-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		out = twoSumSorted(nums, i+1, N-1, -nums[i], out)
	}
	return
}

// Find all "two sum" configurations given a sorted list
// Avoids duplicates
func twoSumSorted(nums []int, lo, hi, target int, out [][]int) [][]int {
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum < target {
			lo += 1
		} else if sum > target {
			hi -= 1
		} else {
			// KEY TRICK:
			//
			// Unlike two sum make sure that we do *all* sums -
			// so when a match happens first skip through duplicates on both ends and then
			// increment lo and decrement hi to continue
			out = append(out, []int{-target, nums[lo], nums[hi]})

			for lo < hi && nums[lo] == nums[lo+1] {
				lo += 1
			}

			for lo < hi && nums[hi] == nums[hi-1] {
				hi -= 1
			}
			// This is needed so that we ensure we are truly skipped past the duplicates
			lo += 1
			hi -= 1
		}
	}
	return out
}

func TestProblem15(t *testing.T) {
	assert.Equal(t, threeSum([]int{1, -1, -1, 0}), [][]int{{-1, 0, 1}})
	assert.Equal(t, threeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{{-1, -1, 2}, {-1, 0, 1}})
	assert.Empty(t, threeSum([]int{0, 1, 1}))
	assert.Equal(t, threeSum([]int{0, 0, 0}), [][]int{{0, 0, 0}})
	assert.Equal(t, threeSum([]int{0, 0, 0, 0}), [][]int{{0, 0, 0}})
	assert.Equal(t, threeSum([]int{-2, 0, 0, 2, 2}), [][]int{{-2, 0, 2}})
}
