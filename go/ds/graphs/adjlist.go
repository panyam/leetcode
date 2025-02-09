package graphs

import "iter"

// Graph representation based on adjacency lists
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
	if !e.Contains(dest) {
		e.dests = append(e.dests, dest)
		e.datas = append(e.datas, data)
	}
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
