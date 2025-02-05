package main

import "github.com/panyam/leetcode/go/ds"

// All the types provided by LeetCode so we dont have to import it
// This is in the same folder so that it can be included within the package without any imports

type ListNode = ds.ListNode[int]

func makeListWithCycle(nth int, values ...int) (*ListNode, *ListNode) {
	head, tail := ds.NewList(false, values...)
	curr := head
	for i := 0; i <= nth; i++ {
		tail.Next = curr
	}
	return head, tail
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Child *Node // For problem 430
}

func NewTree(values ...int) *Node {
	return nil
}
