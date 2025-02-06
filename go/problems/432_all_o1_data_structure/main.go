/*
ProblemId: 432
ProblemTitle: All O`one Data Structure
ProblemLink: https://leetcode.com/problems/merge-two-sorted-lists/
*/
package main

import (
	"container/list"
	"log"
	"maps"
	"slices"
	"strings"
)

type Row struct {
	Count int
	Keys  map[string]bool
}

func NewRow(c int) *Row {
	return &Row{Count: c, Keys: make(map[string]bool)}
}

func (r *Row) Any() string {
	if r == nil {
		return ""
	}
	for k := range r.Keys {
		return k
	}
	return ""
}

func (r *Row) Add(word string) {
	r.Keys[word] = true
}

func (r *Row) Remove(word string) bool {
	if _, ok := r.Keys[word]; ok {
		delete(r.Keys, word)
	}
	return len(r.Keys) > 0
}

type AllOne struct {
	rowmap map[string]*list.Element
	rows   *list.List
}

func (this *AllOne) Print(msg string) {
	log.Println(msg, "-----------")
	for r := this.rows.Front(); r != nil; r = r.Next() {
		row := r.Value.(*Row)
		log.Printf("Count: %d, Keys: %s", row.Count, strings.Join(slices.Collect(maps.Keys(row.Keys)), ", "))
	}
}

// Gets the row corresponding the count == 1
func (this *AllOne) R1() *list.Element {
	head := this.rows.Front()
	if head == nil || head.Value.(*Row).Count > 1 {
		nn := NewRow(1)
		this.rows.PushFront(nn)
	}
	return this.rows.Front()
}

func Constructor() AllOne {
	return AllOne{
		rowmap: map[string]*list.Element{},
		rows:   list.New(),
	}
}

func (this *AllOne) Inc(key string) {
	rowel := this.rowmap[key]
	if rowel == nil {
		rowel = this.R1()
		rowel.Value.(*Row).Add(key)
		this.rowmap[key] = rowel
		return
	}
	row := rowel.Value.(*Row)
	nextRowEl := rowel.Next()
	if nextRowEl == nil || nextRowEl.Value.(*Row).Count != row.Count+1 {
		nn := NewRow(row.Count + 1)
		nextRowEl = this.rows.InsertAfter(nn, rowel)
	}

	nextRow := nextRowEl.Value.(*Row)
	nextRow.Add(key)
	if !row.Remove(key) {
		// ermove row
		this.rows.Remove(rowel)
	}
	this.rowmap[key] = nextRowEl
}

func (this *AllOne) Dec(key string) {
	rowel := this.rowmap[key]
	if rowel == nil {
		// nothing to do
		return
	}

	row := rowel.Value.(*Row)
	nextCount := row.Count - 1
	this.rowmap[key] = nil
	if nextCount > 0 {
		prevRowEl := rowel.Prev()
		if prevRowEl == nil || prevRowEl.Value.(*Row).Count != nextCount {
			nn := NewRow(nextCount)
			prevRowEl = this.rows.InsertBefore(nn, rowel)
		}

		prevRow := prevRowEl.Value.(*Row)
		prevRow.Add(key)
		this.rowmap[key] = prevRowEl
	}
	if !row.Remove(key) {
		// ermove row
		this.rows.Remove(rowel)
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.rows.Back() == nil {
		return ""
	}
	return this.rows.Back().Value.(*Row).Any()
}

func (this *AllOne) GetMinKey() string {
	if this.rows.Front() == nil {
		return ""
	}
	return this.rows.Front().Value.(*Row).Any()
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
