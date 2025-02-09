/*
ProblemId: 3341
ProblemTitle: Minimum time to reach last room
ProblemLink: https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/
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

func minTimeToReach(moveTime [][]int) (out int) {
	R, C := len(moveTime), len(moveTime[0])
	rc2i := func(r, c int) int {
		return r*C + c
	}
	i2rc := func(i int) (r, c int) {
		return int(i / C), i % C
	}

	type Data struct {
		T int
	}

	type Edge = graphs.WeightedEdge[int, int, Data]
	neighbors := func(node int, data Data) iter.Seq[Edge] {
		return func(yield func(Edge) bool) {
			r, c := i2rc(node)
			for _, delta := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				dr, dc := delta[0], delta[1]
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < R && nc >= 0 && nc < C {
					next := rc2i(nr, nc)
					cost := 1 + max(moveTime[nr][nc]-data.T, 0)
					edge := Edge{
						Dest: next,
						Cost: cost,
						Data: Data{
							data.T + cost,
						},
					}
					if !yield(edge) {
						return
					}
				}
			}
		}
	}

	d := graphs.Dijkstra[int, int, Data]{
		Neighbors: neighbors,
		DataLess: func(d1, d2 Data) bool {
			return d1.T < d2.T
		},
	}
	dest := rc2i(R-1, C-1)
	last, _ := d.Run(0, dest)
	return last.Data.T
}
