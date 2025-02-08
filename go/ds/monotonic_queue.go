package ds

// A dequeue where items are only enqueued if it is the smallest item so far.
// Great for problems like 239 - Sliding window max

import "container/list"

type MonotonicQueue struct {
	list.List
	WindowSize int
	Less       func(index1, index2 int) bool
}

func (m *MonotonicQueue) Indexes() (out []int) {
	for c := m.Front(); c != nil; c = c.Next() {
		out = append(out, c.Value.(int))
	}
	return
}

func (m *MonotonicQueue) MaxIndex() int {
	if m.List.Len() == 0 {
		return -1
	}
	return m.Front().Value.(int)
}

func (m *MonotonicQueue) PushIndex(index int) {
	// Since our queue is always sorted (in decreasing order)
	// keep removing items from the LEFT that are SMALLER than the new value
	for m.Len() > 0 {
		frontIndex := m.Front().Value.(int)
		if !m.Less(frontIndex, index) {
			break
		}
		m.Remove(m.Front())
	}

	// Since our queue is always sorted (in decreasing order)
	// keep removing items from the RIGHT that are LARGER than the new value
	for m.Len() > 0 {
		backIndex := m.Back().Value.(int)
		if !m.Less(backIndex, index) {
			break
		}
		m.Remove(m.Back())
	}

	if m.WindowSize > 0 {
		// Trim until window size is honored
		// ie index - Front  <= WindowSize
		for m.Len() > 0 && (index-m.Front().Value.(int) >= m.WindowSize) {
			m.Remove(m.Front())
		}
	}
	m.PushBack(index)
}
