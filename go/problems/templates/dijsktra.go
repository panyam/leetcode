package main

import (
	"container/heap"
	"iter"
)

// Use this template for problems you want to submit that need dijsktra
// Idea is to have all imports "inlined" so copy/paste is easier
// Goal is we should not have any links to internal imports here

////////////// BeginSolution //////////////////
////////////// Example: Problem 3342 //////////////////

func minTimeToReach(moveTime [][]int) (out int) {
	R, C := len(moveTime), len(moveTime[0])
	rc2i := func(r, c int) int {
		return r*C + c
	}
	i2rc := func(i int) (r, c int) {
		return int(i / C), i % C
	}

	type Data struct {
		ArrivalAt     int
		DepartureCost int
	}

	type Edge = WeightedEdge[int, int, Data]
	neighbors := func(node int, data Data) iter.Seq[Edge] {
		return func(yield func(Edge) bool) {
			r, c := i2rc(node)
			// log.Printf("From (%d, A=%d, D=%d) -> ", node, data.ArrivalAt, data.DepartureCost)
			for _, delta := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				dr, dc := delta[0], delta[1]
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < R && nc >= 0 && nc < C {
					next := rc2i(nr, nc)
					cost := 1 + data.DepartureCost + max(moveTime[nr][nc], data.ArrivalAt) - data.ArrivalAt
					edge := Edge{
						Dest: next,
						Cost: cost,
						Data: Data{
							ArrivalAt:     data.ArrivalAt + cost,
							DepartureCost: (data.DepartureCost + 1) % 2,
						},
					}
					// log.Printf("    (%d, A: %d, D: %d)", next, edge.Data.ArrivalAt, edge.Data.DepartureCost)
					if !yield(edge) {
						return
					}
				}
			}
		}
	}

	// log.Println("=============== Starting: ", moveTime)
	d := Dijkstra[int, int, Data]{
		Neighbors: neighbors,
		DataLess: func(d1, d2 Data) bool {
			if d1.ArrivalAt == d2.ArrivalAt {
				return d1.DepartureCost < d2.DepartureCost
			}
			return d1.ArrivalAt < d2.ArrivalAt
		},
	}
	dest := rc2i(R-1, C-1)
	last, _ := d.Run(0, dest)
	return last.Data.ArrivalAt
}

////////////// EndSolution //////////////////

// ///////////// BeginTemplate: go/ds/pq.go //////////////
type PQ[V any] struct {
	items    []V
	LessFunc func(a, b V) bool
}

func (p *PQ[V]) Swap(i, j int) {
	p.items[i], p.items[j] = p.items[j], p.items[i]
}

func (p PQ[V]) Len() int           { return len(p.items) }
func (p PQ[V]) Less(i, j int) bool { return p.LessFunc(p.items[i], p.items[j]) }
func (p *PQ[V]) Push(x any) {
	p.items = append(p.items, x.(V))
}
func (p *PQ[V]) Pop() any {
	n := len(p.items)
	x := p.items[n-1]
	p.items = p.items[0 : n-1]
	return x
}

////////////// EndTemplate: go/ds/pq.go //////////////

// /////////// BeginTemplate: go/ds/graphs/dijsktra.go //////////////
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
	q := PQ[*Edge]{
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
		// log.Println("Next V: ", minEdge)

		// if it has already been visited - we can ignore it as we would ahve a found a closer path through it
		if seen[minEdge.Dest] {
			continue
		}

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
	return
}

///////////// End Template: go/ds/graphs/dijsktra.go //////////////
