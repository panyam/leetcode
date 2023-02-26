
from sortedcontainers import SortedList
from collections import defaultdict

class Comparable:
    """ A way to wrap an arbitrary object with our comparable function. """
    def __init__(self, val, ltfunc):
        self.val = val
        self.ltfunc = ltfunc

    def __lt__(self, another):
        return self.ltfunc(self.val, another.val)

    def __eq__(self, another):
        return self.val == another.val

    def __hash__(self):
        return hash(self.val)

def shortest_path(source, target, neighfunc, INFINITY=10000000000):
    def ltfunc(a, b):
        va = vb = INFINITY
        if a in distances:
            va = distances[a]
        if b in distances:
            vb = distances[b]
        return va < vb

    distances = {source: 0, target: INFINITY}

    # Keeps track of the parent node for a node in the path between source and target.
    parents = {source: None}

    # Which level was a node encountered in
    levels = {source: 0}

    # The nodes that are known to have the shortest path in each iteration
    known_nodes = set([source])

    # A heap of nodes is used where they nodes are sorted by
    # their distance to the source node
    nodeheap = SortedList()

    # Add start's neighbours to heap
    for neighbour,neigh_dist in neighfunc(source, 0, None):
        neigh = Comparable(neighbour, ltfunc)
        distances[neighbour] = neigh_dist
        levels[neighbour] = 1
        parents[neighbour] = source
        nodeheap.add(neigh)

    last = None
    while (last is None or last.val != target) and nodeheap:
        # get the node that is closest to the source at this point
        currnode = nodeheap[0]
        nodeheap.remove(currnode)

        # If the node is already known then skip its children
        if currnode.val in known_nodes: continue

        # Go through each of the curr node's neighbours
        # and update it's distances.   It's "new" distance can either
        # be "via" its parent (currnode) or directly from the
        # source node if such an edge exists.
        currlevel = levels[currnode.val]
        currparent = parents[currnode.val]
        for child,child_dist in neighfunc(currnode.val, currlevel, currparent):
            child = Comparable(child, ltfunc)
            curr_dist = distances[currnode.val] + child_dist
            if child.val not in distances or curr_dist < distances[child.val]:
                distances[child.val] = curr_dist
                levels[child.val] = currlevel + 1
                parents[child.val] = currnode

            # Ensure the heap update's the child priority
            if child in nodeheap:
                nodeheap.remove(child)
            nodeheap.add(child)
        last = currnode
        known_nodes.add(currnode.val)

    # Return the list of all parent nodes that can be walked up
    # backwards to extract the path to the source (in reverse)
    return distances, parents
