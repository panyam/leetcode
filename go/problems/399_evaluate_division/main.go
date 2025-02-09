/*
ProblemId: 399
ProblemTitle: Evaluate Division
ProblemLink: https://leetcode.com/problems/evaluate-division/
*/
package main

import (
	"github.com/panyam/leetcode/go/ds/graphs"
)

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

func makeGraph(equations [][]string, values []float64) *graphs.AdjList[NodeType, EdgeType] {
	adjList := graphs.NewAdjList[NodeType, EdgeType]()
	for i, eqn := range equations {
		edge := values[i]
		adjList.AddEdge(eqn[0], eqn[1], edge)
		adjList.AddEdge(eqn[1], eqn[0], 1/edge)
	}
	return adjList
}

func calcEquation(equations [][]string, values []float64, queries [][]string) (out []float64) {
	graph := makeGraph(equations, values)

	// log.Println("Graph: ", equations)
	for _, q := range queries {
		startVar := q[0]
		endVar := q[1]
		if startVar == endVar {
			if graph.HasVertex(startVar) {
				out = append(out, 1)
			} else {
				out = append(out, -1)
			}
		} else {
			// log.Println("Finding from: ", startVar, endVar)
			sofars := make(map[NodeType]map[NodeType]float64)
			found := false
			bfs := graphs.BFS[NodeType, EdgeType]{Neighbors: graph.Neighbors}
			bfs.EnteringVertex = func(v NodeType) bool {
				if sofars[v] == nil {
					sofars[v] = map[NodeType]float64{
						v: 1,
					}
				}
				return true
			}
			bfs.ProcessEdge = func(fromVar NodeType, weight EdgeType, toVar NodeType) bool {
				sofar := sofars[startVar][fromVar] * weight
				sofars[startVar][toVar] = sofar
				if endVar == toVar {
					found = true
					out = append(out, sofar)
					return false
				}
				return true
			}
			bfs.Run(startVar, true)
			if !found {
				out = append(out, -1)
			}
		}
	}
	return
}
