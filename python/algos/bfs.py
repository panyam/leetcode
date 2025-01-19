def bfs(neighbors_func, start, parents=None, yield_edges=False):
    """ Performs a BFS traversal of a graph from a given node.
    We expect the graph to contain the following:

    neighbors_func: (node) => [child nodes]
        - Returns the neighbors of a given node.

    Some details are:
    1. Does a direct BFS.  
    2. Yields both nodes and edges as traversed and adds "path"
    to a parents list which is finally returned.

    Some times we may want to pass data that is a side car for the node but
    not necessarily a node identifier - this is passed by the startdata obj
    """
    UNDISCOVERED = 0
    DISCOVERED = 1
    VISITED = 2

    state = {start: DISCOVERED}
    queue = [start]
    if parents is None: parents = {}
    parents[start] = None
    level = -1
    while queue:
        level += 1
        newqueue = []
        for currNode in queue:
            if yield_edges:
                result = yield "node", (currNode, level)
            else:
                result = yield (currNode, level)
            if result is False:
                # Then ignore the children for this node!
                continue

            for childNode in neighbors_func(currNode, parents):
                # Process the edge between node and neighbour
                if yield_edges:
                    result = yield "edge", (currNode, childNode)
                    if result is False:
                        return

                if state.get(childNode, UNDISCOVERED) == UNDISCOVERED:
                    """ Only add this if node is not yet discovered. """
                    state[childNode] = DISCOVERED
                    parents[childNode] = currNode
                    newqueue.append(childNode)
            state[currNode] = VISITED
        queue = newqueue
    return parents
