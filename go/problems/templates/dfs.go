package main

import "iter"

// Use this template for problems you want to submit that need dijsktra
// Idea is to have all imports "inlined" so copy/paste is easier
// Goal is we should not have any links to internal imports here

////////////// BeginSolution //////////////////
////////////// Example: Problem 3342 //////////////////

////////////// EndSolution //////////////////

// /////////// BeginTemplate: go/ds/graphs/dfs.go //////////////

// A very simple DFS suitable for unweighted graphs
func SimpleDFS[V comparable](curr V, visited map[V]bool, neighbors func(node V) iter.Seq[V]) iter.Seq[V] {
	if visited == nil {
		visited = map[V]bool{}
	}
	return func(yield func(V) bool) {
		if visited[curr] {
			return
		}
		visited[curr] = true
		defer func() {
			visited[curr] = false
		}()
		if !yield(curr) {
			return
		}
		for next := range neighbors(curr) {
			for yielded := range SimpleDFS(next, visited, neighbors) {
				if !yield(yielded) {
					return
				}
			}
		}
	}
}

type DFS[V comparable, E any] struct {
	Directed   bool
	Processed  map[V]bool
	Discovered map[V]bool
	Parents    map[V]V
	EntryTimes map[V]int
	ExitTimes  map[V]int
	T          int
	YieldEdges bool

	// Method to return the edges
	Edges func(node V) iter.Seq2[V, E]
}

const (
	DFSEntered = 0
	DFSEdge    = 1
	DFSExited  = 2
)

type DFSEvent[V comparable, E any] struct {
	Event int
	Curr  V
	Next  V
	Edge  E
}

func (d *DFS[V, E]) Init() *DFS[V, E] {
	d.Discovered = map[V]bool{}
	d.Processed = map[V]bool{}
	d.Parents = map[V]V{}
	d.EntryTimes = map[V]int{}
	d.ExitTimes = map[V]int{}
	d.T = 0
	return d
}

// A iterator version of the DFS
func (d *DFS[V, E]) Run(curr V) iter.Seq[DFSEvent[V, E]] {
	return func(yield func(DFSEvent[V, E]) bool) {
		d.Discovered[curr] = true
		d.T++

		if !yield(DFSEvent[V, E]{Event: DFSEntered, Curr: curr}) {
			return
		}

		for child, edge := range d.Edges(curr) {
			if !d.Discovered[child] {
				d.Parents[child] = curr

				// Process the edge if need be
				if d.YieldEdges && !yield(DFSEvent[V, E]{Event: DFSEdge, Curr: curr, Next: child, Edge: edge}) {
					return
				}

				// recurse into the children
				d.Run(child)
			} else if d.Directed || !d.Processed[child] {
				if d.YieldEdges && !yield(DFSEvent[V, E]{Event: DFSEdge, Curr: curr, Next: child, Edge: edge}) {
					return
				}
			}
		}

		if !yield(DFSEvent[V, E]{Event: DFSExited, Curr: curr}) {
			return
		}

		d.T++
		d.ExitTimes[curr] = d.T
		d.Processed[curr] = true
	}
}

/*
		Tells if the edge x -> y is such that y is "higher" in the graph
	  than x but we are cycling back.
*/
func (d *DFS[V, E]) IsBackEdge(x, y V) bool {
	return d.Discovered[y] && !d.Processed[y]
}

// Returns true if x is the parent of y
func (d *DFS[V, E]) IsParentEdge(x, y V) bool {
	return d.Parents[y] == x
}

func (d *DFS[V, E]) IsForwardEdge(x, y V) bool {
	return d.Processed[y] && d.EntryTimes[y] > d.EntryTimes[x]
}

func (d *DFS[V, E]) IsCrossEdge(x, y V) bool {
	return d.Processed[y] && d.EntryTimes[y] < d.EntryTimes[x]
}

// TopoLogical sorting that uses DFS above
const (
	TopoSortStarted     = 0
	TopoSortAddedVertex = 1
	TopoSortFoundCycle  = 2
)

type TopoSortEvent[V comparable, E any] struct {
	Event int
	Curr  V
	Next  V
	Edge  E
}

type TopoSort[V comparable, E any] struct {
	Edges  func(node V) iter.Seq2[V, E]
	Output []V
}

func (t *TopoSort[V, E]) Run(nodes []V) iter.Seq[TopoSortEvent[V, E]] {
	return func(yield func(TopoSortEvent[V, E]) bool) {
		dfs := DFS[V, E]{Edges: t.Edges, Directed: true}
		for _, n := range nodes {
			if dfs.Discovered[n] {
				continue
			}

			// Signals start of a new DFS run from a root
			// Indicating a new component
			if !yield(TopoSortEvent[V, E]{Event: TopoSortStarted, Curr: n}) {
				return
			}

			for evt := range dfs.Run(n) {
				if evt.Event == DFSExited {
					t.Output = append(t.Output, evt.Curr)
					if !yield(TopoSortEvent[V, E]{Event: TopoSortAddedVertex, Curr: evt.Curr, Next: evt.Next, Edge: evt.Edge}) {
						return
					}
				} else if evt.Event == DFSEdge {
					curr, next := evt.Curr, evt.Next
					if dfs.IsBackEdge(curr, next) {
						// we have a cycle
						if !yield(TopoSortEvent[V, E]{Event: TopoSortFoundCycle, Curr: curr, Next: next, Edge: evt.Edge}) {
							return
						}
						break
					}
				}
			}
		}
	}
}

///////////// End Template: go/ds/graphs/dijsktra.go //////////////
