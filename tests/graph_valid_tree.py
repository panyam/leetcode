
from tests.utils import run_cases
from algos.grid import *
from algos.dfs import DFS, simpledfs, toposort
from typing import List
import unittest

cases = [
    (
        5,
        [[0,1],[0,2],[0,3],[1,4]],
        True
    ),
    (
        5,
        [[0,1],[1,2],[2,3],[1,3],[1,4]],
        False
    ),
    (
        3,
        [[0,1],[0,2],[1,2]],
        False
    )
]

class TestMethods(unittest.TestCase):
    def test_it(self):
        # And now in dfs
        run_cases(self, cases, solution)

def solution(n: int, edges: List[List[int]]) -> bool:
    neighs = [set() for i in range(n)]
    for n1,n2 in edges:
        neighs[n1].add(n2)
        neighs[n2].add(n1)

    def neighfunc(n, parent=None):
        return neighs[n]

    dfs = DFS(yield_edges=True)
    for _,(src,dst) in filter(lambda evt:evt[0]=="edge", dfs(neighfunc, 0)):
        # if we are not a parent edge then we have a cycle
        # or are a DAG
        print("Edge: ", src, dst)
        if dfs.parents[dst] != src and dfs.parents[src] != dst:
            return False

    for i in range(n):
        if not dfs.discovered[i]: 
            return False
    return True

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
