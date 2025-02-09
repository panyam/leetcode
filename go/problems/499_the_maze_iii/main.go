/*
ProblemId: 499
ProblemTitle: The Maze 3
ProblemLink: https://leetcode.com/problems/the-maze-iii
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
func findShortestWay(maze [][]int, ball []int, hole []int) (instructions string) {
	instructions = "impossible"
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

	sr, sc := ball[0], ball[1]
	startPos := rc2i(sr, sc)

	holeR, holeC := hole[0], hole[1]
	holePos := rc2i(holeR, holeC)

	neighbors := func(pos int, data string) iter.Seq[graphs.WeightedEdge[int, int, string]] {
		dirs := []string{"d", "l", "r", "u"}
		r, c := i2rc(pos)
		// log.Println("r, c: ", R, C, pos, r, c)
		return func(yield func(graphs.WeightedEdge[int, int, string]) bool) {
			// instead of looking at immediate neighbors, go all the way to the end or to a hole
			for dir, delta := range [][]int{{1, 0}, {0, -1}, {0, 1}, {-1, 0}} { // ordered by Down, Left, Right, Up
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
					if cr == holeR && cc == holeC {
						break
					}
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
				cost := max(distR, distC)
				nextPos := rc2i(cr, cc)
				edge := graphs.WeightedEdge[int, int, string]{
					Dest: nextPos,
					Cost: cost,
					Data: data + dirs[dir],
				}
				// log.Printf("Yielding: (%d,%d) -> (%d,%d), Cost: %d", r, c, cr, cc, cost)
				if !yield(edge) {
					return
				}
			}
		}
	}

	log.Printf("(%d,%d) -> (%d, %d)", sr, sc, holeR, holeC)
	d := graphs.Dijkstra[int, int, string]{
		Neighbors: neighbors,
		DataLess: func(d1, d2 string) bool {
			return d1 < d2
		},
	}
	last, found := d.Run(startPos, holePos)
	if !found {
		return
	}
	return last.Data
}
