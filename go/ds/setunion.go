package ds

type SetUnion[V comparable] struct {
	Items   map[V]bool
	Parents map[V]V
	Sizes   map[V]int
	Count   int
}

func (s SetUnion[V]) Init() *SetUnion[V] {
	s.Items = map[V]bool{}
	s.Parents = map[V]V{}
	s.Sizes = map[V]int{}
	s.Count = 0
	return &s
}

func (s *SetUnion[V]) Add(items ...V) *SetUnion[V] {
	for _, item := range items {
		if !s.Has(item) {
			s.Items[item] = true
			s.Parents[item] = item
			s.Sizes[item] = 1
			s.Count += 1
		}
	}
	return s
}

func (s *SetUnion[V]) Has(item V) bool {
	return s.Items[item]
}

// Find the root of a given item.
func (s *SetUnion[V]) Find(item V) V {
	p := s.Parents[item]
	if p == item {
		return p
	}
	return s.Find(p)
}

// Merges two item sets.
func (s *SetUnion[V]) Union(item1, item2 V) {
	// Identify roots of the two items
	p1, p2 := s.Find(item1), s.Find(item2)
	if p1 != p2 {
		s1, s2 := s.Sizes[p1], s.Sizes[p2]
		// Make the "smaller" root a child of the "larger" root
		if s1 >= s2 {
			s.Sizes[p1] = s1 + s2
			s.Parents[p2] = p1
		} else {
			s.Sizes[p2] = s1 + s2
			s.Parents[p1] = p2
		}
	}
}

func (s *SetUnion[V]) SameComp(item1, item2 V) bool {
	return s.Find(item1) == s.Find(item2)
}

/*
   def byparent(s):
       out = {}
       for item in s.Parents.keys():
           parent = s.find(item)
           if parent not in out: out[parent] = []
           out[parent].append(item)
       return out
*/
