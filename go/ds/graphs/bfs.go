package graphs

import "iter"

/*
* Performs a BFS traversal of an UNWEIGHTED graph from a given starting node.
*
* Source: Skiena - Page 164
*
* Useful notes:
*
* - Parents is filled during an interation but can also be passed in with partial results.
*   It allows us to find intersting paths through a graph.   Parents[i] is the vertex that
*   discovered vertex i.
*
*   Except for the root each vertex is discovered during the traversal.  This forms an interesting
*   "discovery tree" with the unique property that the root -> node path is the smallest number of
*   edges.
 */
type BFSIter[V comparable] struct {
	Directed   bool
	Processed  map[V]bool
	Discovered map[V]bool
	Parents    map[V]V
	Neighbors  func(node V) iter.Seq[V]
}

func (b *BFSIter[V]) Run(start V, handler func(event int, level int, currVertex, nextVertex V) int) {
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
		handler = func(event int, level int, currVertex, nextVertex V) int {
			return 0
		}
	}
	// Ensure start has no parents first
	if _, ok := b.Parents[start]; ok {
		delete(b.Parents, start)
	}
	queue := []V{start}
	b.Discovered[start] = true
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
			// log.Println("Curr: ", currVertex)
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

type BFS[V comparable] struct {
	Directed   bool
	Processed  map[V]bool
	Discovered map[V]bool
	Parents    map[V]V
	Neighbors  func(node V) iter.Seq[V]

	// Handlers when nodes and edges are encountered
	EnteringVertex func(dist int, v V) bool
	LeavingVertex  func(dist int, v V) bool
	ProcessEdge    func(dist int, start V, end V) bool
}

func (b *BFS[V]) Run(start V) {
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
	b.Discovered[start] = true
	for level := 0; len(queue) > 0; level++ {
		var nextq []V
		for _, currVertex := range queue {
			if b.EnteringVertex != nil && !b.EnteringVertex(level, currVertex) {
				return
			}

			b.Processed[currVertex] = true

			// Now go through all its children
			// log.Println("Curr: ", currVertex)
			for destVertex := range b.Neighbors(currVertex) {
				// log.Println("Dest: ", destVertex, b.Processed[destVertex])
				if !b.Processed[destVertex] || b.Directed {
					if b.ProcessEdge != nil && !b.ProcessEdge(level, currVertex, destVertex) {
						// we were asked to stop
						return
					}
				}
				if !b.Discovered[destVertex] {
					// Add n to the next queue
					nextq = append(nextq, destVertex)
					b.Discovered[destVertex] = true
					b.Parents[destVertex] = currVertex
				}
			}
			if b.LeavingVertex != nil && !b.LeavingVertex(level, currVertex) {
				return
			}
		}
		// log.Println("Curr, Next Q: ", nextq)
		queue = nextq
	}
}
