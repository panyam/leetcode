/*
ProblemId: 505
ProblemTitle: The Maze 2
ProblemLink: https://leetcode.com/problems/the-maze-ii
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
func shortestDistance(maze [][]int, start []int, destination []int) (dist int) {
	dist = -1
	if len(maze) == 0 || len(maze[0]) == 0 {
		return
	}
	R, C := len(maze), len(maze[0])

	rc2i := func(r, c int) int {
		return r*C + c
	}
	i2rc := func(i int) (r, c int) {
		r = i / C
		c = i % C
		return
	}

	neighbors := func(pos int, data any) iter.Seq[graphs.WeightedEdge[int, int, any]] {
		r, c := i2rc(pos)
		// log.Println("r, c: ", R, C, pos, r, c)
		return func(yield func(graphs.WeightedEdge[int, int, any]) bool) {
			// instead of looking at immediate neighbors, go all the way to the end
			for _, delta := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				dr, dc := delta[0], delta[1]
				cr, cc := r, c
				for true {
					// go as far as we can in this direction
					nr, nc := cr+dr, cc+dc
					if nr < 0 || nc < 0 || nr >= R || nc >= C || maze[nr][nc] == 1 {
						// next stop is a wall so cr,cc is a candidate
						break
					}
					cr, cc = nr, nc
				}
				// yield it
				if cr == r && cc == c {
					continue
				}
				distR, distC := cr-r, cc-c
				if distR < 0 {
					distR = -distR
				}
				if distC < 0 {
					distC = -distC
				}
				nextPos := rc2i(cr, cc)
				cost := max(distR, distC)
				edge := graphs.WeightedEdge[int, int, any]{Dest: nextPos, Cost: cost, Data: any(nil)}
				if !yield(edge) {
					return
				}
			}
		}
	}

	sr, sc := start[0], start[1]
	startPos := rc2i(sr, sc)

	dr, dc := destination[0], destination[1]
	destPos := rc2i(dr, dc)

	// log.Printf("(%d,%d) -> (%d, %d)", sr, sc, dr, dc)
	d := graphs.Dijkstra[int, int, any]{Neighbors: neighbors}
	last, found := d.Run(startPos, destPos)
	if !found {
		return -1
	}
	return last.Cost
}
