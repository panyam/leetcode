
class Node:
    def __init__(self, val):
        self.val = val
        self.next = self.prev = None

def llvalues(node):
    out = []
    curr = node
    while curr:
        out.append(curr.val)
        curr = curr.next
        if curr in (node, None): break
    return out

def llcircmake(*values):
    h1,t1 = llmake(*values)
    if h1 and t1:
        t1.next = h1
        h1.prev = t1
    return h1,t1

def llmake(*values):
    h1 = t1 = None
    for v in values:
        node = Node(v)
        h1,t1 = llinsert(h1, t1, node, t1, doubly=True)
    return h1,t1

def llcat(h1, t1, h2, t2, double=False):
    if not h1:
        return h2, t2
    elif not h2:
        return h1, t1
    else:
        t1.next = h2
        h2.prev = t1
        return h1, t2

def listlen(node):
    """ Return the length of a linkedlist.  """
    i = 0
    while node:
        node = node.next
        i += 1
    return i


def listcmp(l1, l2, cmpfunc=None):
    """ Compares two lists and return -1, 0 or 1 depending on whether
    values of first list <, == or > than values in the second list.

    A custom comparater func can also be provided to which the values of 
    the two nodes can be passed for comparison.
    """
    if not cmpfunc:
        cmpfunc = lambda x,y: x - y

    while l1 and l2:
        diff = cmpfunc(l1.val, l2.val)
        if diff != 0: return diff
        l1 = l1.next
        l2 = l2.next
    if not l1 and not l2: return 0
    elif not l1: return -1
    else: return 1

def lltail(node):
    """ Find the last node in a linked list.
        Note that this also works if our list is circular!
    """
    prev = None
    curr = node
    while curr:
        prev = curr
        curr = curr.next
        if curr == node: break  # will happen if a circular list
    return prev

def llmin(node):
    """
    Find the node with the minimum value in the list.
    Note that this also works if our list is circular!
    """
    min = node
    curr = node
    while curr:
        if curr.val < min.val: min = curr
        curr = curr.next
        if curr == node: break # means we are circular
    return min

def llmid(node):
    """ Find the middle node of a linked (or double linked) list even if the list is circular.
    middle index is len(list) // 2
    """
    prev = None
    slow = fast = node

    while slow and fast:
        # Moving fast first means we dont do any "advance" jumps
        fast = fast.next
        if fast in (node, None): break # 'node' here also checks for circular lists
        else:
            fast = fast.next

        prev = slow
        slow = slow.next

        # Check here if fast is ended otherwise dont bother going back again
        # Especially useful for circular lists
        if fast in (node, None): break
    return prev, slow
    
def llfind(head, findfunc):
    """ Given a head node, finds the first node that matches the criteria from the findfunc.

    Findfunc:
        takes 2 parameters: the prevnode and the current node being tested
        returns the node which maches the criteria along with the previous node (used for insertion)
        If no such node matches than None,None are returned.

    This method will work for circular lists too.
    """
    prev = None
    curr = head
    while curr:
        if findfunc(prev, curr):
            # found the node before which we have to insert
            break
        prev = curr
        curr = curr.next
        if curr == head: 
            # we are back to head after hitting the circular list
            # caller should check if curr == head
            break
    return prev, curr

def lldel(head, tail, node, prev=None, doubly=False):
    if not head: return None, None

    if node is head:
        next = head.next
        head.next = None
        if next:
            if doubly: next.prev = None
            return next, tail
        else:
            return None, None

    next = node.next
    if doubly: prev = node.prev

    assert prev, "Either must be a doubly linked list or a prev must be provided"
    prev.next = next
    if doubly and next: next.prev = prev

    if node is tail: tail = prev
    return head, tail

def llinsert(head, tail, newnode, after=None, doubly=False):
    """
    Insert a newnode after a given 'after' node.  Our list is denoted by the head
    and tail nodes.  Head and/or tail can be null so this method takes care of
    returning the new head/tail after the insertion is complete.  Also all next
    (and prev if doubly linked list) pointers are updated and returned.

    If the after node is null then the item is prepended to the list (and a
    new head will be returned).

    NOTE - the node must *not* be in this list.  newnode's next and prev are first
    clobbered to None so if newnode belongs to this list remove it first with lldel
    """
    newnode.next = None
    if doubly: newnode.prev = None

    if not head:
        return newnode, newnode

    if not after:
        # Prepend
        newnode.next = head
        if doubly: head.prev = newnode
        return newnode, tail

    # Save prev/next pointers
    newnode.next = after.next
    if doubly and newnode.next: newnode.next.prev = newnode

    after.next = newnode
    if doubly:
        newnode.prev = after

    if after == tail: return head, newnode
    else: return head, tail
