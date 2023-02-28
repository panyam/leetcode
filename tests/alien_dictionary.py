
from tests.utils import run_cases
from algos.grid import *
from algos.dfs import DFS, simpledfs, toposort
from typing import List
import unittest

def solution(words: List[str]) -> str:
    from collections import defaultdict
    lessthan = defaultdict(set)
    nodes = set()

    def analyze(w1, w2):
        L1, L2 = len(w1), len(w2)
        for i in range(min(L1, L2)):
            ch1, ch2 = w1[i],w2[i]
            if ch1 != ch2:
                # we have an instance of ordering indicator
                # see if it is alreayd established and matches
                # if not then we have an inconsistency
                if ch1 in lessthan[ch2]:
                    # we have ch2 < ch1 so this cannot be valid
                    return False
                lessthan[ch1].add(ch2)
                return True
        # To handle cases where abc is before ab 
        # (which is not lexicographical sorting)
        return L1 <= L2

cases = [
    (
        ["abc", "ab"], ""
    ),
    (
        ["ab","adc"],
        "abcd"
    ),
    (
        ["wrt","wrf","er","ett","rftt"],
        "wertf"
    ),
    (
        ["z","x"],
        "zx"        
    ),
    (
        ["z","x","z"],
        ""
    ),
]

class TestMethods(unittest.TestCase):
    def test_it(self):
        # And now in dfs
        run_cases(self, cases, solution)

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
