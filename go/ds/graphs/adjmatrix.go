package graphs

import (
	"iter"
)

// Adjacency matrices
// A map of maps with either weights
// Great for dence graphs
type AdjMatrix[N comparable, E any] struct {
	// Weightw can be either 0 or 1 for unweighted graphs or actual values for weighted graphs
	// This means we *cannot* have multi edges but self edges are possible
	EdgeData map[N]map[N]E
}

func NewAdjMatrix[N comparable, E any]() *AdjMatrix[N, E] {
	return &AdjMatrix[N, E]{
		EdgeData: make(map[N]map[N]E),
	}
}

func (a *AdjMatrix[N, E]) Neighbors(node N) iter.Seq[N] {
	return func(yield func(N) bool) {
		tomap := a.EdgeData[node]
		for dest := range tomap {
			if node == dest {
				continue
			}
			if !yield(dest) {
				return
			}
		}
	}
}

func (a *AdjMatrix[N, E]) Edges(node N) iter.Seq2[N, E] {
	return func(yield func(N, E) bool) {
		for k, v := range a.EdgeData[node] {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (a *AdjMatrix[N, E]) AddEdge(src N, dest N, edge E) {
	if a.EdgeData[src] == nil {
		a.EdgeData[src] = map[N]E{}
	}
	if a.EdgeData[dest] == nil {
		a.EdgeData[dest] = map[N]E{}
	}
	a.EdgeData[src][dest] = edge
}

func (a *AdjMatrix[N, E]) Contains(src N) bool {
	return a.EdgeData[src] != nil
}

func (a *AdjMatrix[N, E]) GetEdge(src N, dest N, defaultVal E) E {
	if a.EdgeData[src] == nil {
		return defaultVal
	}
	if out, ok := a.EdgeData[src][dest]; ok {
		return out
	}
	return defaultVal
}
