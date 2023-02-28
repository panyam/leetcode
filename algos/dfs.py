
def simpledfs(neigh, node, visited=None):
    if visited is None: visited = {}

    if not visited.get(node,False):
        visited[node] = True
        yield node
        for nextnode in neigh(node):
            yield from simpledfs(neigh, nextnode, visited)
        visited[node] = False

def toposort(neigh, nodes, output) -> bool:
    dfs = DFS(yield_edges=True)
    for n in nodes:
        if dfs.discovered[n]: continue
        for evttype, evtdata in dfs(neigh, n):
            if evttype == "nodeexited":
                currnode = evtdata
                output.append(currnode)
            elif evttype == "edge":
                currnode,nextnode = evtdata
                if dfs.is_back_edge(currnode, nextnode):
                    return False
    return True

class DFS:
    def __init__(self, directed=False, yield_edges=False, parents=None):
        from collections import defaultdict
        if parents is None: parents = defaultdict(lambda: None)
        self.parents = parents
        self.directed = directed
        self.processed = defaultdict(bool)
        self.discovered = defaultdict(bool)
        self.entry_times = defaultdict(int)
        self.exit_times = defaultdict(int)
        self.yield_edges=yield_edges
        self.T = 0

    def __call__(self, neighbors, currnode):
        """ Performs a DFS traversal of a graph.
        We expect our graph to implement the same interface above. """
        parents = self.parents
        self.discovered[currnode] = True
        self.T += 1
        self.entry_times[currnode] = self.T

        yield "nodeentered", currnode

        for childnode in neighbors(currnode):
            if not self.discovered[childnode]:
                parents[childnode] = currnode

                if self.yield_edges:
                    yield "edge", (currnode, childnode)
                yield from self(neighbors, childnode)
            elif not self.processed[childnode] or self.directed:
                if self.yield_edges: 
                    yield "edge", (currnode, childnode)

        yield "nodeexited", currnode
        
        self.T += 1
        self.exit_times[currnode] = self.T
        self.processed[currnode] = True


    def is_parent_edge(self, x, y):
        """ Returns true if x is the parent of y """
        return self.parents[y] == x

    def is_back_edge(self, x, y):
        """
        Tells if the edge x -> y is such that y is "higher" in the graph
        than x but we are cycling back.
        """
        return self.discovered[y] and not self.processed[y]

    def is_forward_edge(self, x, y):
        return self.processed[y] and self.entry_times[y] > self.entry_times[x]

    def is_cross_edge(self, x,y):
        return self.processed[y] and self.entry_times[y] < self.entry_times[x]
