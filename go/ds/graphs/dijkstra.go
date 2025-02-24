package graphs

import (
	"container/heap"
	"iter"

	"github.com/panyam/leetcode/go/ds"
)

// type EdgeCostExp interface {
// constraints.Integer | constraints.Float
// }

type EdgeCost interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type WeightedEdge[V comparable, E EdgeCost, D any] struct {
	Dest   V
	Cost   E
	Data   D
	Parent *WeightedEdge[V, E, D]
}

type Dijkstra[V comparable, E EdgeCost, D any] struct {
	Parents   map[V]*WeightedEdge[V, E, D]
	Neighbors func(node V, data D) iter.Seq[WeightedEdge[V, E, D]]
	DataLess  func(d1 D, e2 D) bool
}

func (d *Dijkstra[V, E, D]) Run(start, end V) (last *WeightedEdge[V, E, D], found bool) {
	type Edge = WeightedEdge[V, E, D]

	// Our PQ of nodes closest to start
	q := ds.PQ[*Edge]{
		LessFunc: func(e1, e2 *Edge) bool {
			if e1.Cost == e2.Cost && d.DataLess != nil {
				return d.DataLess(e1.Data, e2.Data)
			}
			return e1.Cost < e2.Cost
		},
	}
	startEdge := &Edge{Dest: start, Cost: 0}
	heap.Push(&q, startEdge)

	edgeTo, seen := map[V]*Edge{start: startEdge}, map[V]bool{}
	for q.Len() > 0 {
		// Get the vertex that is closest than all the ones seen so far
		minEdge := heap.Pop(&q).(*Edge)

		// if it has already been visited - we can ignore it as we would ahve a found a closer path through it
		if !seen[minEdge.Dest] {
			seen[minEdge.Dest] = true
			// Add the new vertex to that path
			if minEdge.Dest == end { // if reached the end then return
				return minEdge, true
			}

			for nextEdge := range d.Neighbors(minEdge.Dest, minEdge.Data) { // go through all edges from this vert
				if seen[nextEdge.Dest] { // only consider unseen next-nodes
					continue
				}

				costToNextVertex := minEdge.Cost + nextEdge.Cost
				edgeToDest := edgeTo[nextEdge.Dest]
				if edgeToDest == nil || // Node never seem so add it
					costToNextVertex < edgeToDest.Cost || // This path cost is lower so add it
					(costToNextVertex == edgeToDest.Cost && // cost is same but data is lower cost so add it
						d.DataLess != nil &&
						d.DataLess(nextEdge.Data, edgeToDest.Data)) {
					newEdge := &Edge{Dest: nextEdge.Dest, Cost: costToNextVertex, Data: nextEdge.Data, Parent: minEdge}
					edgeTo[nextEdge.Dest] = newEdge
					heap.Push(&q, newEdge)
				}
			}
		}
	}
	return
}
