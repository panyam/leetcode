
from collections import defaultdict

def simpledfs(neigh, node, data, visited=None):
    if visited is None:
        visited = {}

    if not visited.get(node,False):
        visited[node] = True
        yield node, data
        for nextnode, nextdata in neigh(node, data, None):
            yield from simpledfs(neigh, nextnode, nextdata, visited)
        visited[node] = False

class DFS:
    def __init__(self, is_directed=False):
        self.is_directed = is_directed
        self.processed = defaultdict(bool)
        self.discovered = defaultdict(bool)
        self.entry_times = defaultdict(int)
        self.exit_times = defaultdict(int)
        self.T = 0

    def __call__(self, neighbors, currkey, currdata=None, parents = None):
        """ Performs a DFS traversal of a graph.
        We expect our graph to implement the same interface above. """
        if parents is None: parents = {}

        self.discovered[currkey] = True
        self.T += 1
        self.entry_times[currkey] = self.T

        yield "nodeentered", currkey, currdata

        for childkey, childdata in neighbors(currkey, currdata, parents):
            if not self.discovered[childkey]:
                parents[childkey] = currkey
                yield "edge", currkey, (childkey, childdata)
                yield from self(neighbors, childkey, childdata, parents)
            elif not self.processed[childkey] or self.is_directed:
                yield "edge", currkey, (childkey, childdata)

        yield "nodeexited", currkey, currdata
        
        self.T += 1
        self.exit_times[currkey] = self.T
        self.processed[currkey] = True

