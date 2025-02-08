/*
ProblemId: 55
ProblemTitle: Jump Game
ProblemLink: https://leetcode.com/problems/jump-game/
*/
package main

import "log"

func canJump(nums []int) bool {
	N := len(nums)
	cache := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		v := nums[i]
		if v >= (N-1)-i {
			cache[i] = 1
		} else if v == 0 {
			cache[i] = -1
		} else {
			cache[i] = 0
		}
	}
	return canJump2(nums, 0, cache) > 0
}

func canJump2(nums []int, offset int, cache []int) int {
	N := len(nums)
	if offset == N-1 {
		return 1
	}
	if nums[offset] == 0 {
		return -1
	}
	if cache[offset] == 0 {
		cache[offset] = -1
		for i := 1; i <= nums[offset]; i++ {
			res := canJump2(nums, i+offset, cache)
			cache[offset] = res
			if res > 0 {
				break
			}
			if res == 0 {
				log.Fatalf("how cna res be 0")
			}
		}
	}
	return cache[offset]
}
