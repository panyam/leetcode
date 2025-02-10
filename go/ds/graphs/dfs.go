package graphs

import "iter"

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
				if evt.Event == DFSExitedVertex {
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

// DFS over unweighted graphs
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
	DFSEnteredVertex = 0
	DFSEdge          = 1
	DFSExitedVertex  = 2
)

type DFSEvent[V comparable, E any] struct {
	Event int
	Curr  V
	Next  V
	Edge  E
}

// A iterator version of the DFS
func (d *DFS[V, E]) Run(curr V) iter.Seq[DFSEvent[V, E]] {
	return func(yield func(DFSEvent[V, E]) bool) {
		d.Discovered[curr] = true
		d.T++

		if !yield(DFSEvent[V, E]{Event: DFSEnteredVertex, Curr: curr}) {
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

		if !yield(DFSEvent[V, E]{Event: DFSExitedVertex, Curr: curr}) {
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
