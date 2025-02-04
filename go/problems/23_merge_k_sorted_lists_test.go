/*
ProblemId: 21
ProblemTitle: Merge Two Sorted Lists
ProblemLink: https://leetcode.com/problems/merge-two-sorted-lists/description/
*/
package main

import (
	"container/heap"
	"log"
	"sort"
	"testing"

	"github.com/panyam/leetcode/go/ds"
	"github.com/stretchr/testify/assert"
)

func TestProblem23(t *testing.T) {
	runTest := func(vallists ...[]int) {
		var heads []*ListNode
		var expected []int
		total := 0
		for _, l := range vallists {
			head, _ := ds.NewList(false, l...)
			heads = append(heads, head)
			expected = append(expected, l...)
			total += len(l)
		}
		sort.Ints(expected)

		merged := mergeKLists(heads)
		assert.Equal(t, len(merged.Values()), total)
		assert.EqualValues(t, expected, merged.Values())
	}

	runTest([]int{1, 2, 4}, []int{1, 3, 4})
	runTest()
	runTest(nil)
	runTest(nil, nil)
	runTest(nil, []int{0})
	runTest([]int{1, 4, 5}, []int{1, 3, 6}, []int{2, 6})
}

type LLQueue []*ListNode

func (l LLQueue) Len() int {
	return len(l)
}

func (l LLQueue) Less(i, j int) bool {
	la := l[i]
	lb := l[j]
	if la == nil {
		return lb != nil
	}
	if lb == nil {
		return false
	}
	return la.Val < lb.Val
}

func (l LLQueue) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *LLQueue) Push(x any) {
	// n := len(*l)
	item := x.(*ListNode)
	// item.index = n
	*l = append(*l, item)
}

func (l *LLQueue) Pop() any {
	old := *l
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	// item.index = -1
	*l = old[0 : n-1]
	return item
}

func mergeKLists(heads []*ListNode) *ListNode {
	log.Println("Running: ", heads)
	var lists LLQueue
	heap.Init(&lists)
	for _, h := range heads {
		if h != nil {
			heap.Push(&lists, h)
		}
	}

	var head, tail *ListNode
	for lists.Len() > 0 {
		// log.Println("Lists: ", lists, lists.Len())
		next := heap.Pop(&lists).(*ListNode)
		if head == nil {
			head = next
			tail = next
		} else {
			tail.Next = next
			tail = next
		}
		n2 := next.Next
		next.Next = nil
		if n2 != nil {
			heap.Push(&lists, n2)
		}
	}
	return head
}
