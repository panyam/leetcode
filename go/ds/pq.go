package ds

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
