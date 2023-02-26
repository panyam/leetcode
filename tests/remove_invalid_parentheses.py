from algos.bfs import bfs
from typing import List
import unittest

class TestMethods(unittest.TestCase):
    def test_surrounded_regions(self):
        self.assertEqual(solution("()())()"),  ['(())()', '()()()'])
        self.assertEqual(solution("(a)())()"), ["(a())()","(a)()()"])

# https://leetcode.com/problems/remove-invalid-parentheses/
def solution(s: str) -> List[str]:
    # Our BFS here needs a neighboring func that takes a string and returns
    # a list of strings with 1 "more" paranthesis removed
    def neigh(s):
        for i in range(len(s)):
            if s[i] in ["(", ")"]:
                yield s[:i] + s[i + 1:]

    def isvalid(s):
        count = 0
        for ch in s:
            if ch == "(": count += 1
            if ch == ")": count -= 1
            if count < 0: return False
        return count == 0

    out = []
    for node, level in bfs(neigh, s):
        if isvalid(node):
            if out and len(out[-1]) > len(node):
                break
            out.append(node)
    return out

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
