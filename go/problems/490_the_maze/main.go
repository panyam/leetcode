/*
ProblemId: 490
ProblemTitle: The Maze
ProblemLink: https://leetcode.com/problems/the-maze
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

// Our problem is unweighted, undirected *and* best served with adj-list
func hasPath(maze [][]int, start []int, destination []int) (found bool) {
	R := len(maze)
	if R == 0 {
		return false
	}
	C := len(maze[0])
	if C == 0 {
		return false
	}

	rc2i := func(r, c int) int {
		return r*C + c
	}
	i2rc := func(i int) (r, c int) {
		r = i / C
		c = i % C
		return
	}

	neighbors := func(pos int) iter.Seq[int] {
		r, c := i2rc(pos)
		// log.Println("r, c: ", R, C, pos, r, c)
		return func(yield func(int) bool) {
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
				if !yield(rc2i(cr, cc)) {
					return
				}
			}
		}
	}

	sr, sc := start[0], start[1]
	startPos := rc2i(sr, sc)

	dr, dc := destination[0], destination[1]
	destPos := rc2i(dr, dc)

	log.Printf("(%d,%d) -> (%d, %d)", sr, sc, dr, dc)
	bfs := graphs.BFS[int]{Neighbors: neighbors}
	bfs.EnteringVertex = func(n int) bool {
		if n == destPos {
			found = true
			// check it can indeed stop here
			return false
		}
		return true
	}
	bfs.Run(startPos)
	return
}
