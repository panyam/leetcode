
from algos.grid import *
from algos.bfs import bfs
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
        for case, expected in cases:
            self.assertEqual(solution(case), expected)

# https://leetcode.com/problems/surrounded-regions/
def solution(board: List[List[str]]):
    if not board or not board[0]: return

    R = len(board)
    C = len(board[0])
    nth = 2
    def neigh(loc, data, parents):
        r,c = loc
        for nr,nc in sideneigh(R, C, r, c):
            if board[nr][nc] == "1":
                yield (nr,nc), None

    def solvefrom(r, c):
        for ((nr, nc), level, data) in bfs(neigh, (r,c)):
            board[nr][nc] = nth

    # Solve the border case
    for r in range(R):
        for c in range(C):
            if board[r][c] == "1":
                solvefrom(r, c)
                nth += 1
    return nth - 2

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
