/*
ProblemId: 146
ProblemTitle: LRU Cache
ProblemLink: https://leetcode.com/problems/lru-cache/
*/
package main

import "container/list"

type Item struct {
	Key   int
	Value int
}

type LRUCache struct {
	Capacity int
	Map      map[int]*list.Element
	Elems    *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Elems:    list.New(),
		Map:      make(map[int]*list.Element),
	}
}

func (l *LRUCache) Get(key int) int {
	elem := l.Map[key]
	if elem == nil {
		return -1
	}
	l.Elems.MoveToFront(elem)
	return elem.Value.(*Item).Value
}

func (l *LRUCache) Put(key int, value int) {
	elem := l.Map[key]
	if elem != nil {
		l.Elems.MoveToFront(elem)
		elem.Value.(*Item).Value = value
	} else {
		for l.Elems.Len() >= l.Capacity {
			back := l.Elems.Back()
			l.Elems.Remove(l.Elems.Back())
			l.Map[back.Value.(*Item).Key] = nil
		}
		elem := l.Elems.PushFront(&Item{Key: key, Value: value})
		l.Map[key] = elem
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
