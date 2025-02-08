package graphs

import "iter"

// Adjacency matrices
// A map of maps with either weights
// Great for dence graphs
type AdjMatrix[N comparable, E any] struct {
	// Weightw can be either 0 or 1 for unweighted graphs or actual values for weighted graphs
	// This means we *cannot* have multi edges but self edges are possible
	Weights map[N]map[N]E

	// Whether directed or not
	Directed bool
}

func New[N comparable, E any]() *AdjMatrix[N, E] {
	return &AdjMatrix[N, E]{
		Weights: make(map[N]map[N]E),
	}
}

func (a *AdjMatrix[N, E]) Neighbors(node N) iter.Seq2[N, E] {
	return func(yield func(N, E) bool) {
		for k, v := range a.Weights[node] {
			if !yield(k, v) {
				return
			}
		}
	}
}
