/*
ProblemId: 322
ProblemTitle: Coin Change
ProblemLink: https://leetcode.com/problems/coin-change/
*/
package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BeginProblemTests

func TestProblem322(t *testing.T) {
	runTest := func(coins []int, amount int, expected int) {
		// log.Println("Testing Coins, Amount: ", coins, amount)
		assert.Equal(t, expected, coinChangeIterative(coins, amount))
		assert.Equal(t, expected, coinChangeWithCache(coins, amount, nil))
	}
	runTest([]int{1, 2, 5}, 11, 3)
	runTest([]int{2}, 3, -1)
	runTest([]int{1}, 0, 0)
	runTest([]int{186, 419, 83, 408}, 6249, 20)
	runTest([]int{1, 100}, 2, 2)
}

// EndProblemTests

// Recurrence relation:
// change[amount] = min([ 1 + change(amount - c) for c in coins if amount > c ])
func coinChange(coins []int, amount int) int {
	return coinChangeIterative(coins, amount)
}

// Memoized DB
func coinChangeWithCache(coins []int, amount int, cache map[int]int) int {
	if cache == nil {
		cache = make(map[int]int)
		cache[0] = 0
	}

	if _, ok := cache[amount]; !ok {
		cache[amount] = -1
		for _, c := range coins {
			if amount == c {
				cache[amount] = 1
				break
			} else if amount > c {
				x := coinChangeWithCache(coins, amount-c, cache)
				if x >= 0 && (cache[amount] < 0 || x+1 < cache[amount]) {
					cache[amount] = x + 1
				}
			}
		}
	}
	return cache[amount]
}

// Quick iterative version
func coinChangeIterative(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	cache := make(map[int]int)
	cache[0] = 0
	queue := []int{0}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		var newqueue []int
		for _, v := range queue {
			for i := len(coins) - 1; i >= 0; i-- {
				c := coins[i]
				next := v + c
				if next > amount {
					continue
				}
				val, ok := cache[next]
				if ok {
					cache[next] = min(val, 1+cache[v])
				} else {
					cache[next] = 1 + cache[v]
				}
				if next == amount {
					return cache[next]
				}
				if !visited[next] {
					visited[next] = true
					newqueue = append(newqueue, next)
				}
			}
		}
		queue = newqueue
	}
	return -1
}
