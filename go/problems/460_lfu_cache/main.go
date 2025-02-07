/*
ProblemId: 460
ProblemTitle: LFU Cache
ProblemLink: https://leetcode.com/problems/lfu-cache/
*/
package main

import (
	"container/list"
	"log"
)

type Item struct {
	Key     int
	Value   int
	Counter int
}

type LFUCache struct {
	Capacity int
	Map      map[int]*list.Element
	Elems    *list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Capacity: capacity,
		Elems:    list.New(),
		Map:      make(map[int]*list.Element),
	}
}

func (l *LFUCache) Debug(msg string) {
	log.Println(msg, "==========")
	for n := l.Elems.Front(); n != nil; n = n.Next() {
		log.Println(n.Value)
	}
	log.Println("==========")
}

func (l LFUCache) IncCounter(elem *list.Element) {
	item := elem.Value.(*Item)
	item.Counter++
	// make sure amongst all other items of equal count it is "earlier" in the list to
	for elem.Prev() != nil {
		prev := elem.Prev()
		previtem := prev.Value.(*Item)
		if previtem.Counter <= item.Counter {
			l.Elems.MoveBefore(elem, prev)
		} else {
			break
		}
	}
	// l.Debug("After IncCounter")
}

func (l *LFUCache) Get(key int) int {
	elem := l.Map[key]
	if elem == nil {
		return -1
	}
	l.IncCounter(elem)
	return elem.Value.(*Item).Value
}

func (l *LFUCache) Put(key int, value int) {
	elem := l.Map[key]
	if elem != nil {
		l.IncCounter(elem)
		elem.Value.(*Item).Value = value
	} else {
		// remove the lowest priority item - will be the tail
		if l.Elems.Len() >= l.Capacity {
			back := l.Elems.Back()
			l.Elems.Remove(l.Elems.Back())
			l.Map[back.Value.(*Item).Key] = nil
		}
		elem := l.Elems.PushBack(&Item{Key: key, Value: value, Counter: 0})
		l.IncCounter(elem)
		l.Map[key] = elem
	}
}
