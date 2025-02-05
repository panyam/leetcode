/*
ProblemId: s2296
ProblemTitle: Design a text editor
ProblemLink: https://leetcode.com/problems/merge-two-sorted-lists/
*/
package main

import "log"

// ProblemImpementation:
// Problem impl here fo easy copying
type Node struct {
	Char rune
	Next *Node
	Prev *Node
}

type TextEditor struct {
	// cursor is the current node after which text is appended.
	// it can never be the head node which is used a sentinel
	Dummy *Node
	Curr  *Node
}

func (t *TextEditor) Print(msg string) {
	n := t.Curr
	log.Println(msg, "---------")
	if n == nil {
		log.Println("nil")
		return
	}
	for n.Prev != nil {
		n = n.Prev
	}
	for ; n != nil; n = n.Next {
		if n == t.Dummy {
			log.Print("<DUMMY>")
		} else if n == t.Curr {
			log.Print(string(n.Char), " <- Cursor")
		} else {
			log.Print(string(n.Char))
		}
	}
	log.Println("-----")
}

func Constructor() (out TextEditor) {
	out.Dummy = &Node{Char: 0}
	out.Curr = out.Dummy
	return
}

func toNodes(text string) (h *Node, t *Node) {
	for _, ch := range text {
		n := &Node{Char: ch}
		if h == nil {
			h = n
		} else {
			t.Next = n
			n.Prev = t
		}
		t = n
	}
	return
}

func (this *TextEditor) AddText(text string) {
	h, t := toNodes(text)
	if this.Curr != nil {
		next := this.Curr.Next
		this.Curr.Next = h
		h.Prev = this.Curr
		t.Next = next
		if next != nil {
			next.Prev = t
		}
	}
	this.Curr = t
}

func (this *TextEditor) DeleteText(k int) int {
	i := 0
	for range k {
		if this.Curr == this.Dummy {
			break
		}
		prev := this.Curr.Prev
		next := this.Curr.Next
		if prev != nil {
			prev.Next = next
			this.Curr = prev
		}
		if next != nil {
			next.Prev = prev
		}
		i++
	}
	return i
}

func (this *TextEditor) CursorLeft(k int) string {
	for range k {
		if this.Curr == this.Dummy {
			break
		}
		this.Curr = this.Curr.Prev
	}
	return this.getLast(10)
}

func (this *TextEditor) getLast(rem int) (out string) {
	// Now get the last 10 chars at most
	runes := make([]rune, rem)
	i := rem
	curr := this.Curr
	for range rem {
		if curr == this.Dummy {
			break
		}
		i--
		runes[i] = curr.Char
		curr = curr.Prev
	}
	out = string(runes[i:])
	return
}

func (this *TextEditor) CursorRight(k int) string {
	for range k {
		if this.Curr.Next == nil {
			break
		}
		this.Curr = this.Curr.Next
	}
	return this.getLast(10)
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
