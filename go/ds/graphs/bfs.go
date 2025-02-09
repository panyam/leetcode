package graphs

import "iter"

/*
Performs a BFS traversal of a graph from a given node.
Source: Skiena - Page 164

Useful notes:

  - Parents is filled during an interation but can also be passed in with partial results.
    It allows us to find intersting paths through a graph.   Parents[i] is the vertex that
    discovered vertex i.

    Except for the root each vertex is discovered during the traversal.  This forms an interesting
    "discovery tree" with the unique property that the root -> node path is the smallest number of
    edges.
*/
type BFS[V comparable, E any] struct {
	Processed  map[V]bool
	Discovered map[V]bool
	Parents    map[V]V
	Neighbors  func(node V) iter.Seq2[V, E]

	// Handlers when nodes and edges are encountered
	EnteringVertex func(V) bool
	LeavingVertex  func(V) bool
	ProcessEdge    func(V, E, V) bool
}

func (b *BFS[V, E]) Run(start V, isDirected bool) {
	if b.Processed == nil {
		b.Processed = make(map[V]bool)
	}
	if b.Discovered == nil {
		b.Discovered = make(map[V]bool)
	}
	if b.Parents == nil {
		b.Parents = make(map[V]V)
	}
	// Ensure start has no parents first
	if _, ok := b.Parents[start]; ok {
		delete(b.Parents, start)
	}
	queue := []V{start}
	for len(queue) > 0 {
		var nextq []V
		for _, currVertex := range queue {
			if b.EnteringVertex != nil {
				if !b.EnteringVertex(currVertex) {
					return
				}
			}

			b.Processed[currVertex] = true

			// Now go through all its children
			for n, e := range b.Neighbors(currVertex) {
				if !b.Processed[n] || isDirected {
					if b.ProcessEdge != nil {
						if !b.ProcessEdge(currVertex, e, n) {
							// we were asked to stop
							return
						}
					}
				}
				if !b.Discovered[n] {
					// Add n to the next queue
					nextq = append(nextq, n)
					b.Discovered[n] = true
					b.Parents[n] = currVertex
				}
			}
			if b.LeavingVertex != nil {
				b.LeavingVertex(currVertex)
			}
		}
		queue = nextq
	}
}
