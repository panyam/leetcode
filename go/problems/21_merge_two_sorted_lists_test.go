/*
ProblemId: 21
ProblemTitle: Merge Two Sorted Lists
ProblemLink: https://leetcode.com/problems/merge-two-sorted-lists/description/
*/
package main

import (
	"sort"
	"testing"

	"github.com/panyam/leetcode/go/ds"
	"github.com/stretchr/testify/assert"
)

func TestProblem21(t *testing.T) {
	runTest := func(l1 []int, l2 []int) {
		h1, _ := ds.NewList(false, l1...)
		h2, _ := ds.NewList(false, l2...)
		expected := append(l1, l2...)
		sort.Ints(expected)
		merged := mergeTwoLists(h1, h2)
		assert.Equal(t, len(merged.Values()), len(l1)+len(l2))
		assert.EqualValues(t, expected, merged.Values())
	}

	runTest([]int{1, 2, 4}, []int{1, 3, 4})
	runTest(nil, nil)
	runTest(nil, []int{0})
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var next, head, tail *ListNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			next = list1
			list1 = list1.Next
		} else {
			next = list2
			list2 = list2.Next
		}
		next.Next = nil
		if head == nil {
			head = next
			tail = next
		} else {
			tail.Next = next
			tail = next
		}
	}

	next = nil
	if list1 != nil {
		next = list1
	} else {
		next = list2
	}
	if head == nil {
		head = next
	} else {
		tail.Next = next
	}
	return head
}
