/*
ProblemId: 399
ProblemTitle: Evaluate Division
ProblemLink: https://leetcode.com/problems/evaluate-division/
*/
package main

import (
	"log"

	"github.com/panyam/leetcode/go/ds/graphs"
)

// ProblemImpementation:
// Problem impl here fo easy copying
// For any graph problem:
// 1. first think about graph structure
// - weighted or unweighted
// - directed or undirected
// - adj list or adj matrix

// Our problem is unweighted (it has a value - that does not mean it is weight), directed *and* best served with adj-list
func calcEquation(equations [][]string, values []float64, queries [][]string) (out []float64) {
	adjmatrix := graphs.NewAdjMatrix[string, float64]()

	for i, eqn := range equations {
		a, b := eqn[0], eqn[1]
		adjmatrix.AddEdge(a, b, values[i])
		adjmatrix.AddEdge(b, a, 1/values[i])
	}

	// log.Println("Graph: ", equations)
	for _, q := range queries {
		startVar := q[0]
		endVar := q[1]
		if !adjmatrix.Contains(startVar) || !adjmatrix.Contains(endVar) {
			out = append(out, -1)
		} else if startVar == endVar {
			out = append(out, 1)
		} else {
			log.Println("Finding from: ", startVar, endVar)
			sofars := make(map[string]map[string]float64)
			found := false
			bfs := graphs.BFS[string]{Neighbors: adjmatrix.Neighbors}
			bfs.EnteringVertex = func(v string) bool {
				if sofars[v] == nil {
					sofars[v] = map[string]float64{
						v: 1,
					}
				}
				return true
			}
			bfs.ProcessEdge = func(fromVar string, toVar string) bool {
				weight := adjmatrix.GetEdge(fromVar, toVar, 0)
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
	log.Println("Solution: ", out)
	return
}
