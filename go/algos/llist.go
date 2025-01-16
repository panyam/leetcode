// Single and Double LinkedList implementations
package algos

import "log"

type ListNode[T any] struct {
	Val  T
	Prev *ListNode[T]
	Next *ListNode[T]
}

// Find the tail starting from a given node.
// Even works for circular linked lists
func (n *ListNode[T]) Tail() *ListNode[T] {
	prev := n.Prev
	curr := n
	for curr != nil {
		prev = curr
		curr = curr.Next
		if curr == n {
			break
		}
	}
	return prev
}

// Find the node with the minimum value in the list.
// Note that this also works if our list is circular!
func (n *ListNode[T]) Min(less func(a, b T) bool) (min *ListNode[T], prev *ListNode[T]) {
	curr := n
	min = n
	p := n.Prev
	for curr != nil {
		if less(curr.Val, min.Val) {
			min = curr
			prev = p
		}
		p = curr
		curr = curr.Next
		if curr == n {
			break
		}
	}
	return
}

// Finds the midpoint node starting from this node (as well as its predecessor node).
// Note that this also works if our list is circular!
func (n *ListNode[T]) Mid() (mid *ListNode[T], prev *ListNode[T]) {
	slow := n
	fast := n
	for slow != nil && fast != nil {
		fast = fast.Next
		if fast == nil || fast == n {
			// Also checking for circular lists
			break
		} else {
			fast = fast.Next
		}

		prev = slow
		slow = slow.Next

		// Check here if fast ended otherwise dont bother going back again
		// Especially useful for circular lists
		if fast == n || fast == nil {
			break
		}
	}
	return slow, prev
}

type List[T any] struct {
	Head  *ListNode[T]
	Tail  *ListNode[T]
	Count int
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

// Inserts a node after a particular node
// If pos is nil, then the node is inserted "before" the head node at the front
// Returns the list
func (l *List[T]) Insert(newnode *ListNode[T], after *ListNode[T]) *List[T] {
	// Ensure that newnode does not have a successor or a predecessor
	if newnode.Next != nil || newnode.Prev != nil {
		log.Println("Next or Prev is not nil....")
	}
	newnode.Next = nil
	newnode.Prev = nil
	if l.Head == nil {
		l.Head = newnode
		l.Tail = newnode
	} else if after == nil {
		l.Head.Prev = newnode
		newnode.Next = l.Head
		l.Head = newnode
	} else {
		next := after.Next
		newnode.Next = after.Next

		if next != nil {
			next.Prev = newnode
		}

		after.Next = newnode
		newnode.Prev = after
		if after == l.Tail {
			l.Tail = newnode
		}
	}
	l.Count += 1
	return l
}

// Find a node by a given function
func (l *List[T]) Find(matcher func(val T) bool) (node *ListNode[T], prev *ListNode[T]) {
	curr := l.Head
	for curr != nil {
		if matcher(curr.Val) {
			node = curr
			return node, prev
		}
		prev = curr
		curr = curr.Next
	}
	return
}

// Deletes a particular node.
// Parameter 'prev' is only required if we are talking about singly linked lists.
func (l *List[T]) Delete(node *ListNode[T], prev *ListNode[T]) {
	if l.Head == nil {
		return
	}
	if prev == nil {
		prev = l.Prev(node)
	}
	next := node.Next
	if node == l.Head {
		l.Head = node.Next
	} else if node == l.Tail {
		l.Tail = prev
	}
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}

	l.Count -= 1
	if l.Head == nil || l.Tail == nil {
		l.Tail = nil
		l.Head = nil
		l.Count = 0
	}
}

func (l *List[T]) Prev(node *ListNode[T]) *ListNode[T] {
	if l.Head == nil {
		return nil
	} else if node.Prev != nil || node == l.Head {
		return node.Prev
	}

	// we are a singly LL, so find prev
	curr := l.Head
	prev := l.Head.Prev
	for curr != nil {
		if curr == node {
			return prev
		}
		prev = curr
		curr = curr.Next
	}
	return nil
}
