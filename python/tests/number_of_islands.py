
from algos.grid import *
from algos.bfs import bfs
from algos.dfs import DFS, simpledfs
from typing import List
import unittest

cases = [
    (
        [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]],
        1
    ),
    (
        [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]],
        3
    )
]

class TestMethods(unittest.TestCase):
    def test_surrounded_regions(self):
        # Run all cases in bfs
        for case, expected in cases:
            self.assertEqual(solution(case), expected)

        # And now in dfs
        for case, expected in cases:
            self.assertEqual(solution(case, method="dfs"), expected)

        # And now in simpledfs
        for case, expected in cases:
            self.assertEqual(solution(case, method="simpledfs"), expected)

# https://leetcode.com/problems/surrounded-regions/
def solution(board: List[List[str]], method="bfs"):
    if not board or not board[0]: return

    R = len(board)
    C = len(board[0])
    nth = 2

    def neigh(loc, parents=None):
        r,c = loc
        for nr,nc in sideneigh(R, C, r, c):
            if board[nr][nc] == "1":
                yield (nr,nc)

    def visitfrom(r, c):
        if method == "simpledfs":
            for (nr, nc) in simpledfs(neigh, (r,c)):
                board[nr][nc] = nth
        elif method == "dfs":
            dfs = DFS()
            for evtname, (nr,nc) in dfs(neigh, (r,c)):
                if evtname == "nodeentered":
                    board[nr][nc] = nth
        else:
            for ((nr, nc), level) in bfs(neigh, (r,c)):
                board[nr][nc] = nth

    for (r,c) in filter(lambda x: board[x[0]][x[1]] == "1", itermatrix(R, C)):
        visitfrom(r, c)
        nth += 1

    # clear changes
    for r,c in itermatrix(R, C):
        if board[r][c] != "0": board[r][c] = "1"
    return nth - 2

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
