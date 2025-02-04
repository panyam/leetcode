package main

import (
	"testing"

	"github.com/panyam/leetcode/go/ds"
	"github.com/stretchr/testify/assert"
)

/**
<Problem Number>. <Problem Title> - <Problem Link>

<Problem Description Here>
*/

func runTest(t *testing.T, values ...int) {
	head, _ := ds.NewList(false, values...)
	retvals := reverseList(head).Values()
	lo, hi := 0, len(values)-1
	for lo < hi {
		values[lo], values[hi] = values[hi], values[lo]
		lo += 1
		hi -= 1
	}
	assert.EqualValues(t, values, retvals)
}

func TestProblem206(t *testing.T) {
	runTest(t, 1, 2, 3, 4, 5)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) (out *ListNode) {
	for head != nil {
		next := head.Next
		head.Next = out
		out = head
		head = next
	}
	return
}
