package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem 1: https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	if true { // using sort solution
		indexes := []int{}
		for i := range len(nums) {
			indexes = append(indexes, i)
		}
		sort.Slice(indexes, func(a, b int) bool {
			return nums[indexes[a]] < nums[indexes[b]]
		})
		lo, hi := 0, len(indexes)-1
		for lo < hi {
			s := nums[indexes[lo]] + nums[indexes[hi]]
			if s < target {
				lo += 1
			} else if s > target {
				hi -= 1
			} else {
				return []int{indexes[lo], indexes[hi]}
			}
		}
	} else { // using hashmap solution
		m := make(map[int]int)
		for i := range nums {
			diff := target - nums[i]
			if other, ok := m[diff]; ok {
				return []int{i, other}
			} else {
				m[nums[i]] = i
			}
		}
	}
	return nums
}

func TestTwoSum(t *testing.T) {
	assert.Equal(t, twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
	assert.Equal(t, twoSum([]int{3, 2, 4}, 6), []int{1, 2})
	assert.Equal(t, twoSum([]int{3, 3}, 6), []int{0, 1})
}
