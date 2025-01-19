package main

import "sort"

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
