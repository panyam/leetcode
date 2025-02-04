package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
141. Linked List Cycle

https://leetcode.com/problems/linked-list-cycle/
*/

func TestProblem141(t *testing.T) {
	runTest := func(pos int, values ...int) bool {
		head, _ := makeListWithCycle(pos, values...)
		return hasCycle(head)
	}
	assert.True(t, runTest(1, 3, 2, 0, 4))
	assert.True(t, runTest(0, 1, 2))
	assert.False(t, runTest(-1, 1))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return false
}
