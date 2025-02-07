/*
ProblemId: 460
ProblemTitle: LFU Cache
ProblemLink: https://leetcode.com/problems/lfu-cache/
*/
package main

import (
	"container/list"
	"fmt"
	"log"
)

type Item struct {
	Key     int
	Value   int
	Counter int
	// The Element in the list within the list of values in the Row
	Elem *list.Element
}

type Row struct {
	*list.List
	Counter int
}

func NewRow(counter int) *Row {
	return &Row{
		List:    list.New(),
		Counter: counter,
	}
}

func (r *Row) Add(item *Item) *Item {
	elem := r.PushFront(item)
	item.Counter = r.Counter
	item.Elem = elem
	return item
}

func (r *Row) Remove(item *Item) {
	r.List.Remove(item.Elem)
	item.Elem = nil
}

// Remove the LRU item in this row
func (r *Row) Evict() *Item {
	last := r.Back()
	removed := last.Value.(*Item)
	removed.Elem = nil
	r.List.Remove(last)
	return removed
}

type LFUCache struct {
	Capacity int
	Items    map[int]*Item
	RowList  *list.List
	RowMap   map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	out := LFUCache{
		Capacity: capacity,
		RowList:  list.New(),
		RowMap:   make(map[int]*list.Element),
		Items:    make(map[int]*Item),
	}
	return out
}

func (l *LFUCache) Debug(msg string) {
	log.Println(msg, "==========")
	for n := l.RowList.Front(); n != nil; n = n.Next() {
		row := n.Value.(*Row)
		fmt.Printf("Row [%d]: ", row.Counter)
		for i := row.Front(); i != nil; i = i.Next() {
			item := i.Value.(*Item)
			if item.Counter != row.Counter {
				fmt.Print("\n")
				log.Fatalf("Row: %d/%d, Item (%d, %d) Row (%d) does not match container row (%d)", row.Counter, row.Len(), item.Key, item.Value, item.Counter, row.Counter)
			}
			fmt.Printf("( %d, %d )", item.Key, item.Value)
		}
		fmt.Print("\n")
	}
	log.Println("==========")
}

// Here get the rowElem (from the map) for item.count
// get the nextRowElem for the "next" count
//
// rowElem.Remove(item)
// nextRowElem.Add(item)
//
// if rowElem.Value is empty -> remove rowElem from our list and the map
func (l LFUCache) IncCounter(item *Item) {
	count := item.Counter
	nextCount := count + 1

	rowElem := l.RowMap[count]
	if rowElem == nil {
		log.Fatalf("Row does not exist for count: %d", count)
	}
	nextRowElem := l.RowMap[count+1]
	if nextRowElem == nil {
		// then add one after rowElem
		nextRowElem = l.RowList.InsertAfter(NewRow(nextCount), rowElem)
		l.RowMap[nextCount] = nextRowElem
	}

	row := rowElem.Value.(*Row)
	nextRow := nextRowElem.Value.(*Row)

	row.Remove(item)
	nextRow.Add(item)

	// Is this needed - yes so we can get to the least frequently row in O(1)
	if row.Len() == 0 {
		l.RowList.Remove(rowElem)
		l.RowMap[count] = nil
	}
}

func (l LFUCache) GetR1() *list.Element {
	r1 := l.RowMap[1]
	if r1 == nil {
		r1 = l.RowList.PushFront(NewRow(1))
		l.RowMap[1] = r1
	}
	return r1
}

func (l *LFUCache) Get(key int) int {
	item := l.Items[key]
	if item == nil {
		return -1
	}
	l.IncCounter(item)
	return item.Value
}

func (l *LFUCache) Put(key int, value int) {
	item := l.Items[key]
	if item != nil {
		l.IncCounter(item)
		item.Value = value
	} else {
		// remove the lowest priority item - will be the tail
		l.EnsureSpace()

		newitem := &Item{Key: key, Value: value, Counter: 1}
		row := l.GetR1()
		row.Value.(*Row).Add(newitem)
		l.Items[key] = newitem
	}
}

func (l *LFUCache) EnsureSpace() {
	if len(l.Items) < l.Capacity {
		return
	}

	lowestRowElem := l.RowList.Front()
	lowestRow := lowestRowElem.Value.(*Row)
	removed := lowestRow.Evict()
	if lowestRow.Len() == 0 {
		l.RowList.Remove(lowestRowElem)
		l.RowMap[removed.Counter] = nil
	}
	delete(l.Items, removed.Key)
}
