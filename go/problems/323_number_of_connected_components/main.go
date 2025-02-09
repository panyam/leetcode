/*
ProblemId: 323
ProblemTitle: Number of Connected Components in an Undirected Graph
ProblemLink: https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
*/
package main

import (
	"iter"

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
func countComponents(n int, edges [][]int) (out int) {
	edgelist := map[int][]int{}
	setEdge := func(src, dest int) {
		edgelist[src] = append(edgelist[src], dest)
		edgelist[dest] = append(edgelist[dest], src)
	}
	for _, edge := range edges {
		setEdge(edge[0], edge[1])
	}

	neighbors := func(src int) iter.Seq[int] {
		return func(yield func(int) bool) {
			for _, v := range edgelist[src] {
				if !yield(v) {
					return
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
