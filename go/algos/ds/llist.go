// Single and Double LinkedList implementations
package ds

import "log"

type ListNode[T any] struct {
	Prev *ListNode[T]
	Val  T
	Next *ListNode[T]
}

// Creates a new list with a list of values
func NewListNode[T any](value T) (out *ListNode[T]) {
	return &ListNode[T]{Val: value}
}

func NewList[T any](doubly bool, values ...T) (head, tail *ListNode[T]) {
	for _, v := range values {
		newnode := &ListNode[T]{Val: v}
		head, tail = LNodeInsert(head, tail, newnode, tail, doubly)
	}
	return
}

// Return all values in a list
func (n *ListNode[T]) Values() (out []T) {
	for n != nil {
		out = append(out, n.Val)
		n = n.Next
	}
	return
}

func (n *ListNode[T]) FirstInvalidNode() int {
	var prev *ListNode[T]
	curr := n
	i := 0
	for curr != nil {
		if prev != nil {
			if prev.Next != curr {
				return i
			}
		}
		if curr.Prev != prev {
			return i
		}
		prev = curr
		curr = curr.Next
		i += 1
	}
	return -1
}

// Inserts a node after a given 'after' node.  Our list is denoted by the head and tail nodes.
// Head and/or tail can be null so this method takes care of returning the new head/tail after the
// insertion is complete.  Also all next (and prev if doubly linked) pointers are updated and
// returned.
//
// If the after node is null then the item is prepended to the list (and a new head will
// be returned).
//
// NOTE - newnode must NOT be in this list.  newnode's next and prev are first
// clobbered to nil so if newnode belongs to this, remove it first (with Delete)
func LNodeInsert[T any](head, tail, newnode *ListNode[T], after *ListNode[T], doubly bool) (*ListNode[T], *ListNode[T]) {
	// Ensure that newnode does not have a successor or a predecessor
	if newnode.Next != nil || newnode.Prev != nil {
		log.Println("Next or Prev is not nil....")
	}
	newnode.Next = nil
	if doubly {
		newnode.Prev = nil
	}

	// Empty list
	if head == nil {
		return newnode, newnode
	}

	// Insert at the front
	if after == nil {
		newnode.Next = head
		if doubly {
			head.Prev = newnode
		}
		return newnode, tail
	}

	// Save next/prev pointers
	newnode.Next = after.Next
	after.Next = newnode
	if doubly {
		if newnode.Next != nil {
			newnode.Next.Prev = newnode
		}
		newnode.Prev = after
	}

	if after == tail {
		tail = newnode
	}
	return head, tail
}

// Deletes a particular node.
// Parameter 'prev' is only required if we are talking about singly linked lists.
func LNodeDelete[T any](head, tail, node, prev *ListNode[T], doubly bool) (*ListNode[T], *ListNode[T]) {
	if head == nil {
		return nil, nil
	}

	// Replacing the head - this also takes care of when head == tail
	if node == head {
		next := head.Next

		// Clean up head next/prev
		head.Next = nil

		if next != nil {
			if doubly {
				next.Prev = nil
			}
			return next, tail
		} else {
			return nil, nil
		}
	}

	// Normal case where prev exists
	if doubly {
		prev = node.Prev
	} else if prev == nil || prev.Next != node {
		panic("In a singly linked list prev must be non nil and prev.next == node")
	}

	next := node.Next
	prev.Next = next
	if doubly {
		next.Prev = prev
	}

	if node == tail {
		tail = prev
	}
	return head, tail
}

// Starting at head of a singly linked list, returns the "previous" node for a given node.
// Not really needed for a doubly linked list.
func LNodePrev[T any](head, node *ListNode[T]) (prev *ListNode[T]) {
	if head == nil {
		return nil
	}

	// we are a singly LL, so find prev
	curr := head
	prev = nil
	for curr != nil {
		if curr == node {
			return prev
		}
		prev = curr
		curr = curr.Next
	}
	return nil
}

// cancatenates two linked lists
func LNodeCat[T any](h1, t1, h2, t2 *ListNode[T], doubly bool) (nh, nt *ListNode[T]) {
	if h1 == nil {
		return h2, t2
	}
	if h2 == nil {
		return h1, t1
	}
	t1.Next = h2
	if doubly {
		h2.Prev = t1
	}
	return h1, t2
}

// Compares the list starting at this node with another list starting at 'another' node
// and returns -1, 0 or 1 depending on the comparator func which compares the two values
func (n *ListNode[T]) Compare(another *ListNode[T], cmpfn func(a, b T) int) int {
	l1 := n
	l2 := another
	for l1 != nil && l2 != nil {
		v := cmpfn(l1.Val, l2.Val)
		if v != 0 {
			return v
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil {
		return 0
	} else if l1 == nil {
		return -1
	} else {
		return 1
	}
}

// Finds the length of the list starting from this node.
func (n *ListNode[T]) Length() (l int) {
	for n != nil {
		l += 1
		n = n.Next
	}
	return
}

// Starting at a head node, finds the first node that matches the given matcher function.
// matcher:
// Takes 2 arguments: the prevnode and the currnode being tested and returns true if a match is found.
//
// Returns the node which matches the criteria along with the previous node.
// The previous node is useful for singly linked list incase an insertion has to
// happen at this point by the caller.
func (n *ListNode[T]) Find(matcher func(val T) bool) (prev *ListNode[T], curr *ListNode[T]) {
	curr = n
	for curr != nil {
		if matcher(curr.Val) {
			return prev, curr
		}
		prev = curr
		curr = curr.Next
	}
	return
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
func (n *ListNode[T]) Min(less func(a, b T) bool) (prev *ListNode[T], min *ListNode[T]) {
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
func (n *ListNode[T]) Mid() (prev *ListNode[T], mid *ListNode[T]) {
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
	return prev, slow
}
