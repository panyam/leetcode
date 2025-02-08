package graphs

import "iter"

// In or graphs we abstract how Neighbors are fetched
// As long as the Neighbors function gives us a set of neighbors for a given node we are fine.

// Our graphs need to answer a few questions
type Graph[N comparable, E any] interface {
	// Neighbors is an iterator
	Neighbors(node N) iter.Seq2[N, E]
}
