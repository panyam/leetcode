/*
ProblemId: 847
ProblemTitle: Shortest Path Visiting All Nodes
ProblemLink: https://leetcode.com/problems/shortest-path-visiting-all-nodes/
*/
package main

import (
	"iter"
	"math"

	"github.com/panyam/leetcode/go/ds/graphs"
)

// ProblemImpementation:
// Problem impl here fo easy copying
// For any graph problem:
// - is it weighted or unweighted
// - is it directed or undirected
// - do we need adj list or adj matrix

func shortestPathLength(graph [][]int) (out int) {
	N := len(graph)
	RIGHTBITS := (1 << (N)) - 1
	createNode := func(node, rightbits int) int {
		return node<<(N+1) | rightbits
	}
	splitNode := func(node int) (int, int) {
		return node >> (N + 1), node & RIGHTBITS
	}
	neighbors := func(node int) iter.Seq[int] {
		return func(yield func(int) bool) {
			vertIndex, usedUp := splitNode(node)
			edges := graph[vertIndex]
			// log.Println("Edges: ", vertIndex, strconv.FormatInt(int64(usedUp), 2), edges)
			for _, nextIndex := range edges {
				nextUsedUp := usedUp | (1 << nextIndex)
				nextNode := createNode(nextIndex, nextUsedUp)
				// log.Println("Next Node: ", nextNode, nextIndex, usedUp|(1<<nextIndex))
				// log.Println("Edge: ", strconv.FormatInt(int64(node), 2), " -> ", strconv.FormatInt(int64(nextNode), 2), "userUp: ", strconv.FormatInt(int64(nextUsedUp), 2))
				if !yield(nextNode) {
					return
				}
			}
		}
	}
	out = math.MaxUint32
	for i := range N {
		startNode := createNode(i, 1<<i)
		// log.Println("N: ", N, strconv.FormatInt(int64(RIGHTBITS), 2))
		// log.Println("Staring at: ", strconv.FormatInt(int64(startNode), 2))
		bfs := graphs.BFSIter[int]{Neighbors: neighbors}
		bfs.Run(startNode, func(event int, level int, curr, next int) int {
			if event == -1 {
				_, usedUp := splitNode(curr)
				if usedUp == RIGHTBITS {
					out = min(level, out)
					return -1
				}
				if level > out {
					return 1
				}
			}
			return 0
		})
	}
	return
}
