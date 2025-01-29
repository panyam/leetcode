package main

import (
	"testing"

	"github.com/panyam/leetcode/go/ds"
	"github.com/stretchr/testify/assert"
)

type ListNode = ds.ListNode[int]

// Problem 2: https://leetcode.com/problems/add-two-numbers/description/
//
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (out *ListNode) {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		digit := sum % 10
		carry = sum / 10
		n := ds.NewListNode(digit)
		if head == nil {
			head = n
			tail = n
		} else {
			tail.Next = n
			tail = n
		}
	}
	return head
}

func TestAddTwoNumbers(t *testing.T) {
	h1, _ := ds.NewList(false, 2, 4, 3)
	h2, _ := ds.NewList(false, 5, 6, 4)
	exp, _ := ds.NewList(false, 7, 0, 8)
	assert.Equal(t, addTwoNumbers(h1, h2).Compare(exp, func(a, b int) int { return a - b }), 0)

	h1, _ = ds.NewList(false, 0)
	h2, _ = ds.NewList(false, 0)
	exp, _ = ds.NewList(false, 0)
	assert.Equal(t, addTwoNumbers(h1, h2).Compare(exp, func(a, b int) int { return a - b }), 0)

	h1, _ = ds.NewList(false, 9, 9, 9, 9, 9, 9, 9)
	h2, _ = ds.NewList(false, 9, 9, 9, 9)
	exp, _ = ds.NewList(false, 8, 9, 9, 9, 0, 0, 0, 1)
	assert.Equal(t, addTwoNumbers(h1, h2).Compare(exp, func(a, b int) int { return a - b }), 0)
}
