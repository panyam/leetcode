/*
ProblemId: 3341
ProblemTitle: Minimum time to reach last room
ProblemLink: https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/
*/
package main

import (
	"iter"
	"log"
)

// ProblemImpementation:
// Problem impl here fo easy copying
// For any graph problem:
// - is it weighted or unweighted
// - is it directed or undirected
// - do we need adj list or adj matrix

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) (out []int) {
	type Node struct {
		Index     int
		PathSoFar []int
	}

	RED := 1
	BLUE := -1
	edges := NewAdjList[int, int]()
	for _, edge := range redEdges {
		edges.AddEdge(edge[0]+1, edge[1]+1, RED)
	}
	for _, edge := range blueEdges {
		edges.AddEdge(edge[0]+1, edge[1]+1, BLUE)
	}

	neighbors := func(node int) iter.Seq[int] {
		return func(yield func(int) bool) {
			// came from a red edge, so now only yield blue edges
			// log.Println("Curr: ", node)
			if node == 0 {
				// start state so add everything
				for next, color := range edges.Edges(node + 1) {
					// Color does not matter - just add it
					// came from blue so only use RED edges
					// or
					// came from a blue edge, so only yield red edges
					// log.Println("    Next, Color: ", next, color)
					if !yield(color * next) {
						return
					}
				}
			} else {
				realNode := node
				if realNode < 0 {
					realNode = -realNode
				}
				for next, color := range edges.Edges(realNode) {
					if (node < 0 && color == RED) || (node > 0 && color == BLUE) {
						// log.Println("    Next: ", next*color)
						// came from blue so only use RED edges
						// or
						// came from a blue edge, so only yield red edges
						if !yield(color * next) {
							return
						}
					}
				}
			}
		}
	}

	for range n {
		out = append(out, -1)
	}
	out[0] = 0

	b := BFSIter[int]{Neighbors: neighbors}
	startNodes := []int{0}
	b.Run(startNodes, func(event int, level int, currVertex, nextVertex int) int {
		if event == -1 { // Entering
			realNode := -1
			if currVertex < 0 {
				realNode = (-currVertex) - 1
			} else {
				realNode = currVertex - 1
			}
			log.Println("RealNode: ", realNode, out)
			if realNode > 0 && out[realNode] < 0 {
				out[realNode] = level
			}
		}
		return 0
	})
	return
}

// ///////////// BeginTemplate: go/ds/graph/bfs.go //////////////
// A small variation/extension on BFSIter where instead of taking a single node as start
// we are adding known "start" states in one go - This will avoid us having to add
// dummy sentinel states etc
type BFSIter[V comparable] struct {
	Directed   bool
	Processed  map[V]bool
	Discovered map[V]bool
	Parents    map[V]V
	Neighbors  func(node V) iter.Seq[V]
}

func (b *BFSIter[V]) Run(start []V, handler func(event int, level int, currVertex, nextVertex V) int) {
	if b.Processed == nil {
		b.Processed = make(map[V]bool)
	}
	if b.Discovered == nil {
		b.Discovered = make(map[V]bool)
	}
	if b.Parents == nil {
		b.Parents = make(map[V]V)
	}
	if handler == nil {
		handler = func(event int, level int, currVertex, nextVertex V) int { return 0 }
	}
	// Ensure start has no parents first
	var queue []V
	for _, s := range start {
		if _, ok := b.Parents[s]; ok {
			delete(b.Parents, s)
		}
		queue = append(queue, s)
		b.Discovered[s] = true
	}
	for level := 0; len(queue) > 0; level++ {
		var nextq []V
		for _, currVertex := range queue {
			if res := handler(-1, level, currVertex, currVertex); res < 0 {
				return
			} else if res > 0 {
				continue
			}

			b.Processed[currVertex] = true

			// Now go through all its children
			for destVertex := range b.Neighbors(currVertex) {
				// log.Println("Dest: ", destVertex, b.Processed[destVertex])
				if !b.Processed[destVertex] || b.Directed {
					if res := handler(0, level, currVertex, destVertex); res < 0 {
						return
					} else if res > 0 {
						continue
					}
				}
				if !b.Discovered[destVertex] {
					// Add n to the next queue
					nextq = append(nextq, destVertex)
					b.Discovered[destVertex] = true
					b.Parents[destVertex] = currVertex
				}
			}
			if res := handler(1, level, currVertex, currVertex); res < 0 {
				return
			}
		}
		// log.Println("Curr, Next Q: ", nextq)
		queue = nextq
	}
}

// ///////////// EndTemplate: go/ds/graph/bfs.go //////////////

// ///////////// BeginTemplate: go/ds/graph/adjlist.go //////////////

type AdjList[N comparable, E any] struct {
	// Edges[i] denotes all list of edges for a node i
	edges map[N]*EdgeList[N, E]
}

func NewAdjList[N comparable, E any]() *AdjList[N, E] {
	return &AdjList[N, E]{
		edges: make(map[N]*EdgeList[N, E]),
	}
}

func (a *AdjList[N, E]) HasVertex(node N) bool {
	return a.edges[node] != nil
}

func (a *AdjList[N, E]) Neighbors(node N) iter.Seq[N] {
	return func(yield func(N) bool) {
		edges := a.edges[node]
		if edges.Len() == 0 {
			return
		}
		for _, dest := range edges.dests {
			if !yield(dest) {
				return
			}
		}
	}
}

func (a *AdjList[N, E]) Edges(node N) iter.Seq2[N, E] {
	return func(yield func(N, E) bool) {
		edges := a.edges[node]
		if edges.Len() == 0 {
			return
		}
		for i, dest := range edges.dests {
			if !yield(dest, edges.datas[i]) {
				return
			}
		}
	}
}

func (a *AdjList[N, E]) AddEdge(src N, dest N, edge E) {
	edges := a.edges[src]
	// TODO - check if it already exists?
	if edges == nil {
		edges = &EdgeList[N, E]{}
		a.edges[src] = edges
	}
	edges.Add(dest, edge)
}

// A list of edges
type EdgeList[N comparable, E any] struct {
	dests []N
	datas []E
}

func (e *EdgeList[N, E]) Add(dest N, data E) *EdgeList[N, E] {
	// Do not check for duplicates - leave it to the caller to avoid parallel edges
	e.dests = append(e.dests, dest)
	e.datas = append(e.datas, data)
	return e
}

func (e *EdgeList[N, E]) Len() int {
	if e == nil {
		return 0
	}
	return len(e.dests)
}

func (e *EdgeList[N, E]) Contains(v N) bool {
	for _, dest := range e.dests {
		if dest == v {
			return true
		}
	}
	return false
}

// ///////////// EndTemplate: go/ds/graph/adjlist.go //////////////
