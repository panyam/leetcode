/*
ProblemId: 143
ProblemTitle: Reorder List
ProblemLink: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
*/
package main

import (
	"testing"

	"github.com/panyam/leetcode/go/ds"
	"github.com/stretchr/testify/assert"
)

func TestProblem143(t *testing.T) {
	runTest := func(vals ...int) {
		var expected []int
		lo, hi := 0, len(vals)-1
		for lo <= hi {
			expected = append(expected, vals[lo])
			if lo != hi {
				expected = append(expected, vals[hi])
			}
			lo += 1
			hi -= 1
		}
		head, _ := ds.NewList(false, vals...)

		result := reorderList(head)
		assert.Equal(t, len(expected), len(result.Values()))
		assert.EqualValues(t, expected, result.Values())
	}

	runTest(1, 2, 3, 4)
	runTest(1, 2, 3, 4, 5)
	runTest()
	runTest(1)
	runTest(1, 2)
}

func ListNodeReverse(head *ListNode) (out *ListNode) {
	for head != nil {
		next := head.Next
		head.Next = out
		out = head
		head = next
	}
	return
}

func ListNodeMid(head *ListNode) (prev *ListNode, mid *ListNode) {
	if head == nil || head.Next == nil {
		return nil, head
	}

	mid = head
	fast := head.Next
	for mid != nil && fast != nil {
		prev = mid
		mid = mid.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return
}

func ListNodeAppend(head, tail, newnode *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		head = newnode
		tail = newnode
	} else {
		tail.Next = newnode
		tail = newnode
	}
	tail.Next = nil
	return head, tail
}

func reorderList(head *ListNode) *ListNode {
	prev, mid := ListNodeMid(head)
	if prev == nil || mid == nil {
		return head
	}

	prev.Next = nil

	h1, h2 := head, ListNodeReverse(mid)

	// Now merge them
	// log.Println("H1: ", h1.Values())
	// log.Println("H2: ", h2.Values())

	var h3, t3 *ListNode
	for h1 != nil || h2 != nil {
		if h1 != nil {
			next := h1.Next
			h3, t3 = ListNodeAppend(h3, t3, h1)
			h1 = next
		}
		if h2 != nil {
			next := h2.Next
			h3, t3 = ListNodeAppend(h3, t3, h2)
			h2 = next
		}
	}

	return h3
}
