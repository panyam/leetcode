/*
ProblemId: s2296
ProblemTitle: Design a text editor
ProblemLink: https://leetcode.com/problems/merge-two-sorted-lists/
*/
package main

import (
	"container/list"
	"log"
)

type TextEditor struct {
	// cursor is the current node after which text is appended.
	// it can never be the head node which is used a sentinel
	Dummy *list.Element
	Chars *list.List
	Curr  *list.Element
}

func Constructor() (out TextEditor) {
	out.Chars = list.New()
	out.Dummy = out.Chars.PushFront(0)
	out.Curr = out.Dummy
	return
}

func (t *TextEditor) Print(msg string) {
	log.Println(msg, "---------")
	for n := t.Chars.Front(); n != nil; n = n.Next() {
		if n == t.Dummy {
			log.Print("<DUMMY>")
		} else if n == t.Curr {
			log.Print(n.Value, " <- Cursor")
		} else {
			log.Print(n.Value)
		}
	}

	log.Println("-----")
}

func (this *TextEditor) AddText(text string) {
	for _, ch := range text {
		this.Curr = this.Chars.InsertAfter(ch, this.Curr)
	}
}

func (this *TextEditor) DeleteText(k int) int {
	i := 0
	for range k {
		if this.Curr == this.Dummy {
			break
		}
		prev := this.Curr.Prev()
		this.Chars.Remove(this.Curr)
		this.Curr = prev
		i++
	}
	return i
}

func (this *TextEditor) CursorLeft(k int) string {
	for range k {
		if this.Curr == this.Dummy {
			break
		}
		this.Curr = this.Curr.Prev()
	}
	return this.getLast(10)
}

func (this *TextEditor) CursorRight(k int) string {
	for range k {
		if this.Curr.Next() == nil {
			break
		}
		this.Curr = this.Curr.Next()
	}
	return this.getLast(10)
}

func (this *TextEditor) getLast(rem int) (out string) {
	// Now get the last 10 chars at most
	runes := make([]rune, rem)
	i := rem
	for curr := this.Curr; i > 0 && curr != this.Dummy; curr = curr.Prev() {
		i--
		runes[i] = curr.Value.(rune)
	}
	out = string(runes[i:])
	return
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
