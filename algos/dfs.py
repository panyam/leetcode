
def toposort(neigh, nodes):
    dfs = DFS()
    output = []
    for n in nodes:
        if dfs.discovered(n): continue
        for evttype, currnode in dfs(neigh, n):
            if evttype == "nodeexited":
                output.push(currnode)
            elif evttype == "edge":
                if dfs.is_back_edge(currnode, nextnode):
                    assert False, "have a cycle"
    return output

def simpledfs(neigh, node, visited=None):
    if visited is None: visited = {}

    if not visited.get(node,False):
        visited[node] = True
        yield node
        for nextnode in neigh(node):
            yield from simpledfs(neigh, nextnode, visited)
        visited[node] = False

class DFS:
    def __init__(self, is_directed=False, parents=None):
        from collections import defaultdict
        if parents is None: parents = {}
        self.parents = parents
        self.is_directed = is_directed
        self.processed = defaultdict(bool)
        self.discovered = defaultdict(bool)
        self.entry_times = defaultdict(int)
        self.exit_times = defaultdict(int)
        self.T = 0

    def __call__(self, neighbors, currkey):
        """ Performs a DFS traversal of a graph.
        We expect our graph to implement the same interface above. """
        parents = self.parents
        self.discovered[currkey] = True
        self.T += 1
        self.entry_times[currkey] = self.T

        yield "nodeentered", currkey

        for childkey in neighbors(currkey):
            if not self.discovered[childkey]:
                parents[childkey] = currkey
                yield "edge", currkey, (childkey)
                yield from self(neighbors, childkey)
            elif not self.processed[childkey] or self.is_directed:
                yield "edge", (currkey, childkey)

        yield "nodeexited", currkey
        
        self.T += 1
        self.exit_times[currkey] = self.T
        self.processed[currkey] = True


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
