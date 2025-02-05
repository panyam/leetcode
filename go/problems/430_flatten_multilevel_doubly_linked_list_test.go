/*
ProblemId: 430
ProblemTitle: Flatten a Multilevel Doubly Linked List
ProblemLink: https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/
*/
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BeginProblemTests

// TBD - need a tree construction method
func TestProblem430(t *testing.T) {
	assert.True(t, true)
}

// EndProblemTests

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	h, t := toDoubleList(root)
	if t != nil {
		t.Right = h
		h.Left = t
	}
	return h
}
