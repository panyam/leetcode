/*
ProblemId: 399
ProblemTitle: Evaluate Division
ProblemLink: https://leetcode.com/problems/evaluate-division/
*/
package main

import "github.com/panyam/leetcode/go/ds/graphs"

type EdgeList struct {
}

// ProblemImpementation:
// Problem impl here fo easy copying
// For any graph problem:
// 1. first think about graph structure
// - weighted or unweighted
// - directed or undirected
// - adj list or adj matrix

// Our problem is weighted, directed *and* best served with adj-list
type NodeType = string
type EdgeType = float64

func makeGraph(equations [][]string, values []float64) graphs.Graph[NodeType, EdgeType] {
	var adjList graphs.AdjList[NodeType, EdgeType]
	for i, eqn := range equations {
		edge := values[i]
		adjList.AddEdge(eqn[0], eqn[1], edge)
	}
	return &adjList
}

func calcEquation(equations [][]string, values []float64, queries [][]string) (out []float64) {
	graph := makeGraph(equations, values)

	for _, q := range queries {
		var bfs graphs.BFS[NodeType, EdgeType]
		bfs.Run(q[0], true, graph.Neighbors)
	}
	return
}
