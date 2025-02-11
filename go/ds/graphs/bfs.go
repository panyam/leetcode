package graphs

import "iter"

/*
* Performs a BFS traversal of an UNWEIGHTED graph from a given starting node.  Source: Skiena - Page 164
*
* - Parents is filled during an interation but can also be passed in with partial results.
*   It allows us to find intersting paths through a graph.   Parents[i] is the vertex that
*   discovered vertex i.
*
*   Except for the root each vertex is discovered during the traversal.  This forms an interesting
*   "discovery tree" with the unique property that the root -> node path is the smallest number of
*   edges.
 */
// type VertexType = any
type BFS[VertexType comparable] struct {
	Directed   bool
	Processed  map[VertexType]bool
	Discovered map[VertexType]bool
	Parents    map[VertexType]VertexType
	Neighbors  func(node VertexType) iter.Seq[VertexType]

	// Handlers when nodes and edges are encountered
	EnteringVertex func(dist int, v VertexType) bool
	LeavingVertex  func(dist int, v VertexType) bool
	ProcessEdge    func(dist int, start VertexType, end VertexType) bool
}

func (b *BFS[VertexType]) Init() *BFS[VertexType] {
	if b.Processed == nil {
		b.Processed = make(map[VertexType]bool)
	}
	if b.Discovered == nil {
		b.Discovered = make(map[VertexType]bool)
	}
	if b.Parents == nil {
		b.Parents = make(map[VertexType]VertexType)
	}
	return b
}

func (b *BFS[VertexType]) Run(start VertexType) {
	b.Init()
	// Ensure start has no parents first
	if _, ok := b.Parents[start]; ok {
		delete(b.Parents, start)
	}
	queue := []VertexType{start}
	b.Discovered[start] = true
	for level := 0; len(queue) > 0; level++ {
		var nextq []VertexType
		for _, currVertex := range queue {
			if b.EnteringVertex != nil && !b.EnteringVertex(level, currVertex) {
				return
			}

			b.Processed[currVertex] = true

			// Now go through all its children
			for destVertex := range b.Neighbors(currVertex) {
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
		queue = nextq
	}
}
