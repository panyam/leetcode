
from tests.utils import run_cases
from algos.grid import *
from algos.dfs import DFS, simpledfs, toposort
from typing import List
import unittest

def solution(numCourses: int, prerequisites: List[List[int]]) -> List[int]:
    needs = [set() for n in range(numCourses)]
    for course,dependency in prerequisites:
        needs[course].add(dependency)

    def neighbors(n, parents=None):
        return needs[n]

    output = []
    for evt, data in toposort(neighbors, range(numCourses), output):
        if evt == "CYCLE":
            return []
    return output

cases = [
    (
        2, [[1,0]],
        [0,1]
    ),
    (
        4, [[1,0],[2,0],[3,1],[3,2]],
        [0,1,2,3]
    ),
    (
        1, [] ,
        [0]
    ),
]

class TestMethods(unittest.TestCase):
    def test_it(self):
        # And now in dfs
        run_cases(self, cases, solution)

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
