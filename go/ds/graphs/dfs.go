package graphs

import "iter"

// A more thorough DFS for graphs with custom edges etc based on Skiena
type DFS[V comparable, E any] struct {
	Directed   bool
	Processed  map[V]bool
	Discovered map[V]bool
	Parents    map[V]V
	EntryTimes map[V]int
	ExitTimes  map[V]int
	T          int

	// Method to return the edges
	Edges func(node V) iter.Seq2[V, E]

	// Handler methods
	EnteringVertex func(v V) bool
	LeavingVertex  func(v V) bool
	ProcessEdge    func(start, end V, edge E) bool
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
func (d *DFS[V, E]) Run(curr V) (finished bool) {
	d.Discovered[curr] = true
	d.T++

	defer func() {
		d.T++
		d.ExitTimes[curr] = d.T
		d.Processed[curr] = true
	}()

	if d.EnteringVertex != nil && !d.EnteringVertex(curr) {
		return false
	}

	for child, edge := range d.Edges(curr) {
		if !d.Discovered[child] {
			d.Parents[child] = curr

			// Process the edge if need be
			if d.ProcessEdge != nil && !d.ProcessEdge(curr, child, edge) {
				return false
			}

			// recurse into the children
			d.Run(child)
		} else if d.Directed || !d.Processed[child] {
			if d.ProcessEdge != nil && !d.ProcessEdge(curr, child, edge) {
				return false
			}
		}
	}

	if d.LeavingVertex != nil && !d.LeavingVertex(curr) {
		return false
	}
	return true
}

func (d *DFS[V, E]) IsTreeEdge(x, y V) bool {
	return d.Parents[y] == x
}

func (d *DFS[V, E]) IsBackEdge(x, y V) bool {
	return d.Discovered[y] && !d.Processed[y]
}

func (d *DFS[V, E]) IsForwardEdge(x, y V) bool {
	return d.Processed[y] && d.EntryTimes[y] > d.EntryTimes[x]
}

func (d *DFS[V, E]) IsCrossEdge(x, y V) bool {
	return d.Processed[y] && d.EntryTimes[y] < d.EntryTimes[x]
}

// TopoLogical sorting that uses DFS above
type TopoSort[V comparable, E any] struct {
	Edges  func(node V) iter.Seq2[V, E]
	Output []V

	Started     func(vertex V) bool
	AddedVertex func(vertex V) bool
	FoundCycle  func(curr, next V, edge E) bool
}

func (t *TopoSort[V, E]) Run(nodes []V) (finished bool) {
	dfs := (&DFS[V, E]{Edges: t.Edges, Directed: true}).Init()
	dfs.LeavingVertex = func(v V) bool {
		t.Output = append(t.Output, v)
		if t.AddedVertex != nil && !t.AddedVertex(v) {
			return false
		}
		return true
	}
	dfs.ProcessEdge = func(curr, next V, edge E) bool {
		if dfs.IsBackEdge(curr, next) {
			// we have a cycle
			if t.FoundCycle != nil && !t.FoundCycle(curr, next, edge) {
				return false
			}
		}
		return true
	}

	for _, n := range nodes {
		if !dfs.Discovered[n] {

			// Signals start of a new DFS run from a root
			// Indicating a new component
			if t.Started != nil && !t.Started(n) {
				return false
			}

			dfs.Run(n)
		}
	}
	return true
}

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
