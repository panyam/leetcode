/*
ProblemId: 432
ProblemTitle: All O`one Data Structure
ProblemLink: https://leetcode.com/problems/merge-two-sorted-lists/
*/
package main

import (
	"log"
	"maps"
	"slices"
	"strings"
)

type Row struct {
	Count int
	Keys  map[string]bool
	Next  *Row
	Prev  *Row
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

func (r *Row) AddWord(word string) {
	r.Keys[word] = true
}

func (r *Row) RemoveWord(word string) {
	if _, ok := r.Keys[word]; ok {
		delete(r.Keys, word)
	}
}

type AllOne struct {
	rowmap map[string]*Row
	head   *Row
	tail   *Row
}

func (this *AllOne) Print(msg string) {
	log.Println(msg, "-----------")
	for r := this.head; r != nil; r = r.Next {
		log.Printf("Count: %d, Keys: %s", r.Count, strings.Join(slices.Collect(maps.Keys(r.Keys)), ", "))
	}
}

// Gets the row corresponding the count == 1
func (this *AllOne) getR1() *Row {
	if this.head == nil || this.head.Count > 1 {
		nn := NewRow(1)
		nn.Next = this.head
		if this.head != nil {
			this.head.Prev = nil
		}
		this.head = nn
	}
	if this.head.Next == nil {
		this.tail = this.head
	}
	return this.head
}

func (this *AllOne) getNextRow(r *Row) *Row {
	next := r.Next
	if next == nil || next.Count != r.Count+1 {
		// we are at the tail
		nn := NewRow(r.Count + 1)
		r.Next = nn
		nn.Prev = r
		nn.Next = next
		if next != nil {
			next.Prev = nn
		} else {
			this.tail = nn
		}
	}
	return r.Next
}

func (this *AllOne) getPrevRow(r *Row) *Row {
	prev := r.Prev
	if prev == nil || prev.Count != r.Count-1 {
		// we are at the tail
		pp := NewRow(r.Count - 1)
		r.Prev = pp
		pp.Next = r
		pp.Prev = prev
		if prev != nil {
			prev.Next = pp
		} else {
			this.head = pp
		}
	}
	return r.Prev
}

func (this *AllOne) removeRow(r *Row) {
	this.Print("...")
	log.Print("Before removing row: ", r.Count, r.Keys)
	prev := r.Prev
	next := r.Next
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
	if r == this.head {
		this.head = next
	}
	if r == this.tail {
		this.tail = prev
	}
	if this.head == nil || this.tail == nil {
		this.head = nil
		this.tail = nil
	}
	r.Next = nil
	r.Prev = nil
	log.Print("After removing row: ", r.Count, r.Keys)
	this.Print("...")
}

func Constructor() AllOne {
	return AllOne{
		rowmap: make(map[string]*Row),
	}
}

func (this *AllOne) Inc(key string) {
	row := this.rowmap[key]
	if row == nil {
		// it does not exist - so let us add it to row 1
		row = this.getR1()
		row.AddWord(key)
		this.rowmap[key] = row
	} else {
		nextrow := this.getNextRow(row)
		row.RemoveWord(key)
		if len(row.Keys) == 0 {
			this.removeRow(row)
		}
		nextrow.AddWord(key)
		this.rowmap[key] = nextrow
	}
}

func (this *AllOne) Dec(key string) {
	row := this.rowmap[key]
	if row == nil {
		// doesnt exist so return
		return
	}

	// it does not exist - so let us add it to row 1
	if row.Count > 1 {
		prevrow := this.getPrevRow(row)
		prevrow.AddWord(key)
		this.rowmap[key] = prevrow
	} else {
		this.rowmap[key] = nil
	}

	row.RemoveWord(key)
	if len(row.Keys) == 0 {
		this.removeRow(row)
	}
}

func (this *AllOne) GetMaxKey() (key string) {
	this.Print("Calling GetmaxKey")
	return this.tail.Any()
}

func (this *AllOne) GetMinKey() (key string) {
	return this.head.Any()
}
