/*
ProblemId: 426
ProblemTitle: Convert Binary Search Tree to Sorted Doubly Linked List
ProblemLink: https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/
*/
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BeginProblemTests

// TBD - need a tree construction method
func TestProblem426(t *testing.T) {
	assert.True(t, true)
}

// EndProblemTests

func treeToDoublyList(root *Node) *Node {
	h, t := toDoubleList(root)
	if t != nil {
		t.Right = h
		h.Left = t
	}
	return h
}

func toDoubleList(root *Node) (head *Node, tail *Node) {
	if root == nil {
		return nil, nil
	}
	lh, lt := toDoubleList(root.Left)
	rh, rt := toDoubleList(root.Right)

	if lt != nil {
		head = lh
		lt.Right = root
		root.Left = lt
	} else {
		head = root
	}

	if rh != nil {
		root.Right = rh
		rh.Left = root
		tail = rt
	} else {
		tail = root
	}
	return
}
