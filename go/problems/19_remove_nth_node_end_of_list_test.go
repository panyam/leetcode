/*
ProblemId: 19
ProblemTitle: Remove Nth Node From End of List
ProblemLink: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
*/
package main

import (
	"testing"

	"github.com/panyam/leetcode/go/ds"
	"github.com/stretchr/testify/assert"
)

func TestProblem19(t *testing.T) {
	runTest := func(nth int, vals ...int) {
		var expected []int
		head, _ := ds.NewList(false, vals...)
		for i, v := range vals {
			if i != len(vals)-nth {
				expected = append(expected, v)
			}
		}

		result := removeNthFromEnd(head, nth)
		assert.Equal(t, len(expected), len(result.Values()))
		assert.EqualValues(t, expected, result.Values())
	}

	runTest(2, 1, 2, 3, 4, 5)
	runTest(1, 1)
	runTest(1, 1, 2)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nth := head
	for range n - 1 {
		if nth != nil {
			nth = nth.Next
		}
	}

	// here nth - head == n
	// log.Println("1. n, values: ", n, head.Values())
	if nth == nil {
		return head
	}
	var prev *ListNode
	curr := head
	for nth.Next != nil {
		prev = curr
		curr = curr.Next
		nth = nth.Next
	}

	if prev == nil {
		return head.Next
	}
	if prev.Next != nil {
		prev.Next = prev.Next.Next
	}
	// log.Println("2. n, values: ", n, head.Values())

	return head
}
