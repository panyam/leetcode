package ds

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNodeLen(t *testing.T) {
	h, _ := NewList[int](false)
	fmt.Printf("Length: %d\n", h.Length())
	assert.Equal(t, h.Length(), 0, "Length should be 0")

	h, _ = NewList(false, 1, 2, 3, 4, 5)
	assert.Equal(t, h.Length(), 5, "Length should be 5")
}

func TestInsert(t *testing.T) {
	h1, t1 := NewList[int](true)
	assert.Nil(t, t1)
	assert.Nil(t, h1)

	n := NewListNode(1)
	h2, t2 := LNodeInsert(h1, t1, n, nil, true)
	assert.NotEqual(t, h2, h1)
	assert.Equal(t, h2.Next, h1)
	assert.Equal(t, h2, h2)

	h3, t3 := LNodeInsert(h2, t2, NewListNode(2), nil, true)
	log.Println(h3.Values(), h3, h3.Next, t3)
	assert.NotEqual(t, h2, h3)
	assert.Equal(t, h3.Next, h2)
	assert.Equal(t, h3.Next, t3)
	assert.Equal(t, t2, t3)
	assert.Equal(t, t3.Prev, h3)

	n = NewListNode(3)
	h4, t4 := LNodeInsert(h3, t3, n, t3, true)
	assert.Equal(t, h4.Values(), []int{2, 1, 3})
	assert.Equal(t, h4.FirstInvalidNode(), -1)
	assert.Equal(t, h3, h4)
	assert.NotEqual(t, t3, t4)

	n = NewListNode(4)
	h5, t5 := LNodeInsert(h4, t4, n, t4.Prev, true)
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
	fmt.Printf("Cmp: %d\n", h1.Compare(h2, func(a, b int) int { return a - b }))

	h1, _ = NewList(true, 1, 2, 3, 4, 5)
	h2, _ = NewList(false, 1, 2, 3, 4, 7)
	fmt.Printf("Cmp: %d\n", h1.Compare(h2, func(a, b int) int { return a - b }))
	// Output:
	// Cmp: 0
	// Cmp: -1
	// Cmp: 1
	// Cmp: -2
}

func ExampleTestTail() {
	h1, t1 := NewList(true, 1, 2, 3, 4, 5)
	fmt.Printf("Tail Matches: %t", h1.Tail() == t1)
	// Output:
	// Tail Matches: true
}

func ExampleListMin() {
	h1, t1 := NewList(true, 1, 2, 3, 4, 0)
	_, min := h1.Min(func(a, b int) bool { return a < b })
	fmt.Printf("%t", min == t1)
	// Output:
	// true
}

func _ExampleListMid() {
	h1, t1 := NewList(true, 1)
	prev, mid := h1.Mid()
	fmt.Println(prev == nil)
	fmt.Println(mid == t1)

	h1, t1 = NewList(true, 1, 2)
	prev, mid = h1.Mid()
	fmt.Println(prev == h1)
	fmt.Println(mid == t1)

	h1, t1 = NewList(true, 10, 4, 9, 11, 1, 200)
	prev, mid = h1.Mid()
	log.Println("H1, T1: ", h1, t1)
	fmt.Println(prev == h1.Next.Next)
	fmt.Println(mid == t1.Prev.Prev)
	// Output:
	// true
	// true
	// true
	// true
}

/*
class TestMethods(unittest.TestCase):
    def test_llmid(self):
        h1,t1 = llmake
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (9, 11))

        h1,t1 = llmake(10, 4, 9, 11, 1, 200, 300)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (9, 11))

    def test_llmid_circular(self):
        h1,t1 = llcircmake(1)
        prev,mid = llmid(h1)
        self.assertEqual((prev, mid.val), (None, 1))

        h1,t1 = llcircmake(1, 2)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (1, 2))

        h1,t1 = llcircmake(10, 4, 9, 11, 1, 200)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (9, 11))

        h1,t1 = llcircmake(10, 4, 9, 11, 1, 200, 300)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (9, 11))

    def test_lldel(self):
        # Insert into empty list
        h1, t1 = llmake(1,2,3,4,5)

        # remove head
        h2, t2 = lldel(h1, t1, h1)
        self.assertEqual(llvalues(h2), [2, 3, 4, 5])
        self.assertEqual(h2.val, 2)
        self.assertEqual(t2.val, 5)

        # remove tail
        h3, t3 = lldel(h2, t2, t2, doubly=True)
        self.assertEqual(llvalues(h3), [2, 3, 4])
        self.assertEqual(h3.val, 2)
        self.assertEqual(t3.val, 4)

        # val from the middle
        h4, t4 = lldel(h3, t3, h3.next, doubly=True)
        self.assertEqual(llvalues(h3), [2, 4])
        self.assertEqual(h4.val, 2)
        self.assertEqual(t4.val, 4)
*/
