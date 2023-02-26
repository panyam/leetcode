
from algos.heaps import Comparable
from algos.djikstra import shortest_path 
from typing import List
import unittest
from collections import defaultdict

def solution_sp(n, flights, src, dst, k):
    weights = defaultdict(lambda: defaultdict(int))
    for s,d,cost in flights: weights[s][d] = cost

    def neighbors(city, level, parent):
        if level <= k:
            for city, cost in weights[city].items():
                yield city, cost

    prices, parents = shortest_path(src, dst, neighbors)
    print("Prices: ", prices, parents)
    return -1 if prices[dst] > 100000 else prices[dst]

def solution(n, flights, src, dst, k):
    weights = defaultdict(lambda: defaultdict(int))
    for s,d,cost in flights: weights[s][d] = cost

    def neighbors(city, level, parent):
        if level <= k:
            for city, cost in weights[city].items():
                yield city, cost

    # gen = bfs(neighbors, src)

    prices, parents = shortest_path(src, dst, neighbors)
    print("Prices: ", prices, parents)
    return -1 if prices[dst] > 100000 else prices[dst]


class TestMethods(unittest.TestCase):
    def test_surrounded_regions(self):
        # self.assertEqual(solve_surrounded_regions(grid),solution)
        for case in cases[3:]:
            self.assertEqual(solution(*case["input"]), case["output"])

cases = [
    {
        "input": (
            4, [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]],
            0, 3, 1
        ),
        "output": 700,
    },
    {
        "input": (
            3,
            [[0,1,100],[1,2,100],[0,2,500]],
            0, 2, 1
        ),
        "output": 200,
    },
    {
        "input": (
            3,
            [[0,1,100],[1,2,100],[0,2,500]],
            0, 2, 0,
        ),
        "output": 500,
    },
    {
        "input": (
            4,
            [[0,1,1],[0,2,5],[1,2,1],[2,3,1]],
            0, 3, 1
        ),
        "output": 6,
    }
]

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
