from algos.grid import *
from algos.bfs import bfs
from typing import List
import unittest

class TestMethods(unittest.TestCase):
    def test_surrounded_regions(self):
        grid = [["X","X","X","X"],
                ["X","O","O","X"],
                ["X","X","O","X"],
                ["X","O","X","X"]]
        solution = [["X","X","X","X"],
                    ["X","X","X","X"],
                    ["X","X","X","X"],
                    ["X","O","X","X"]]
        self.assertEqual(solve_surrounded_regions(grid),solution)

        # Case 2
        grid = [["X"]]
        solution = [["X"]]
        self.assertEqual(solve_surrounded_regions(grid), solution)

        # Case 3
        grid = [["0", "0"], ["0", "0"]]
        solution = [["0", "0"], ["0", "0"]]
        self.assertEqual(solve_surrounded_regions(grid), solution)


# https://leetcode.com/problems/surrounded-regions/
def solve_surrounded_regions(board: List[List[str]]):
    if not board or not board[0]: return

    R = len(board)
    C = len(board[0])
    def neigh(loc, parents=None):
        r,c = loc
        for nr,nc in sideneigh(R, C, r, c):
            if board[nr][nc] == "O":
                yield (nr,nc)

    def solve_border(r, c):
        if board[r][c] == 'O':
            for ((nr, nc), level) in bfs(neigh, (r,c)):
                board[nr][nc] = "Y"

    def solve_inland(r, c):
        if board[r][c] == 'O':
            for ((nr, nc), level) in bfs(neigh, (r,c)):
                board[nr][nc] = "X"

    # Solve the border case
    for r in range(R):
        solve_border(r, 0)
        solve_border(r, C - 1)

    for c in range(C):
        solve_border(0, c)
        solve_border(R - 1, c)

    for r in range(R):
        for c in range(C):
            solve_inland(r, c)

    # Fix border case
    for r in range(R):
        for c in range(C):
            if board[r][c] == "Y":
                board[r][c] = "O"
    return board

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
