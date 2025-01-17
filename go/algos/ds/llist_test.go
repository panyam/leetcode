package ds

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNodeLen(t *testing.T) {
	h, _ := NewList[int](false)
	assert.Equal(t, h.Length(), 0, "Length should be 0")

	h, _ = NewList(false, 1, 2, 3, 4, 5)
	assert.Equal(t, h.Length(), 5, "Length should be 5")
}

func TestInsert(t *testing.T) {
	h1, t1 := NewList[int](true)
	assert.Nil(t, t1)
	assert.Nil(t, h1)

	n := NewListNode(1)
	h2, t2 := h1.Insert(t1, n, nil, true)
	assert.NotEqual(t, h2, h1)
	assert.Equal(t, h2.Next, h1)
	assert.Equal(t, h2, h2)

	h3, t3 := h2.Insert(t2, NewListNode(2), nil, true)
	assert.NotEqual(t, h2, h3)
	assert.Equal(t, h3.Next, h2)
	assert.Equal(t, h3.Next, t3)
	assert.Equal(t, t2, t3)
	assert.Equal(t, t3.Prev, h3)

	n = NewListNode(3)
	h4, t4 := h3.Insert(t3, n, t3, true)
	assert.Equal(t, h4.Values(), []int{2, 1, 3})
	assert.Equal(t, h4.FirstInvalidNode(), -1)
	assert.Equal(t, h3, h4)
	assert.NotEqual(t, t3, t4)

	n = NewListNode(4)
	h5, t5 := h4.Insert(t4, n, t4.Prev, true)
	assert.Equal(t, h5.Values(), []int{2, 1, 4, 3})
	assert.Equal(t, h5.FirstInvalidNode(), -1)
	assert.Equal(t, h5, h4)
	assert.NotEqual(t, t5, t)
}

func TestListNodeCmp(t *testing.T) {
	h1, _ := NewList(true, 1, 2, 3, 4, 5)
	h2, _ := NewList(false, 1, 2, 3, 4, 5)
	assert.Equal(t, h1.Compare(h2, func(a, b int) int { return a - b }), 0)

	h1, _ = NewList(true, 1, 2, 3, 4, 5)
	h2, _ = NewList(false, 1, 2, 3, 4, 5, 6)
	assert.Equal(t, h1.Compare(h2, func(a, b int) int { return a - b }), -1)

	h1, _ = NewList(true, 1, 2, 3, 4, 5, 6)
	h2, _ = NewList(false, 1, 2, 3, 4, 5)
	assert.Equal(t, h1.Compare(h2, func(a, b int) int { return a - b }), 1)

	h1, _ = NewList(true, 1, 2, 3, 4, 5)
	h2, _ = NewList(false, 1, 2, 3, 4, 7)
	assert.Equal(t, h1.Compare(h2, func(a, b int) int { return a - b }), -2)
}

func TestListNodeTail(t *testing.T) {
	h1, t1 := NewList(true, 1, 2, 3, 4, 5)
	assert.Equal(t, h1.Tail(), t1)
	assert.Equal(t, t1.Val, 5)
}

func TestListMin(t *testing.T) {
	h1, t1 := NewList(true, 1, 2, 3, 4, 0)
	prev, min := h1.Min(func(a, b int) bool { return a < b })
	assert.Equal(t, prev, t1.Prev)
	assert.Equal(t, min, t1)
}

func TestListMid(t *testing.T) {
	h1, t1 := NewList(true, 9, 11)
	assert.Equal(t, h1.FirstInvalidNode(), -1)
	prev, mid := h1.Mid()
	assert.Equal(t, prev, h1)
	assert.Equal(t, mid, t1)

	h1, t1 = NewList(true, 10, 4, 9, 11, 1, 200)
	assert.Equal(t, h1.FirstInvalidNode(), -1)
	prev, mid = h1.Mid()
	assert.Equal(t, prev, h1.Next.Next)
	assert.Equal(t, mid, t1.Prev.Prev)
}

func TestListDelete(t *testing.T) {
	h1, t1 := NewList(true, 1, 2, 3, 4, 5)
	h2, t2 := h1.Delete(t1, h1, nil, true)
	// test removeing  head
	assert.Equal(t, h2.Values(), []int{2, 3, 4, 5})
	assert.Equal(t, h2.Val, 2)
	assert.Equal(t, t2.Val, 5)

	// test removing tail
	log.Println(h2.Values())
	h3, t3 := h2.Delete(t2, t2, nil, true)
	assert.Equal(t, h3.Values(), []int{2, 3, 4})
	assert.Equal(t, h3.Val, 2)
	assert.Equal(t, t3.Val, 4)

	// removing from the middle
	h4, t4 := h3.Delete(t3, h3.Next, nil, true)
	assert.Equal(t, h3.Values(), []int{2, 4})
	assert.Equal(t, h4.Val, 2)
	assert.Equal(t, t4.Val, 4)
}

func TestCircularList(t *testing.T) {
	h1, t1 := NewCircularList(true, 1)
	prev, mid := h1.Mid()
	assert.Nil(t, prev)
	assert.Equal(t, mid, h1)
	assert.Equal(t, mid, t1)

	h1, t1 = NewCircularList(true, 1, 2)
	prev, mid = h1.Mid()
	assert.Equal(t, prev, h1)
	assert.Equal(t, mid, t1)

	h1, _ = NewCircularList(true, 10, 4, 9, 11, 1, 200)
	prev, mid = h1.Mid()
	assert.Equal(t, prev, h1.Next.Next)
	assert.Equal(t, mid, prev.Next)

	h1, t1 = NewCircularList(true, 10, 4, 9, 11, 1, 200, 300)
	prev, mid = h1.Mid()
	assert.Equal(t, prev, mid.Prev)
	assert.Equal(t, mid, t1.Prev.Prev.Prev)
}
