/*
ProblemId: 547
ProblemTitle: Number of Provinces
ProblemLink: https://leetcode.com/problems/number-of-provinces
*/
package main

import (
	"iter"
	"log"

	"github.com/panyam/leetcode/go/ds/graphs"
)

// ProblemImpementation:
// Problem impl here fo easy copying
// For any graph problem:
// - is it weighted or unweighted
// - is it directed or undirected
// - do we need adj list or adj matrix

// Our problem is WEIGHTED - unlike maze 1.  In maze-1 it did not matter which way we went
// here we could go two diff ways and how far we go - is the length of that edge
// so here dijkstra is needed
func findCircleNum(isConnected [][]int) (out int) {
	n := len(isConnected)
	neighbors := func(src int) iter.Seq[int] {
		return func(yield func(int) bool) {
			log.Println("Neighbors of ", src, isConnected[src])
			for i, v := range isConnected[src] {
				if i != src && v == 1 {
					if !yield(i) {
						return
					}
				}
			}
		}
	}

	b := graphs.BFS[int]{Neighbors: neighbors}
	for i := range n {
		if !b.Processed[i] {
			out += 1
			// start a bfs here
			b.Run(i)
		}
	}
	return
}
