from algos.grid import *

import unittest
class TestMethods(unittest.TestCase):
    def test8(self):
        self.assertEqual(list(allneigh(5, 5, 0, 0)),
                         [(0, 1), (1, 0), (1, 1)])
        self.assertEqual(list(allneigh(5, 5, 4, 4)),
                         [(3, 3), (3, 4), (4, 3)])
        self.assertEqual(list(allneigh(5, 5, 0, 4)),
                         [(0, 3), (1, 3), (1, 4)])
        self.assertEqual(list(allneigh(5, 5, 4, 0)),
                         [(3, 0), (3, 1), (4, 1)])
        self.assertEqual(list(allneigh(5, 5, 2, 2)),
                        [(1, 1), (1, 2), (1, 3), (2, 1),
                         (2, 3), (3, 1), (3, 2), (3, 3)])

    def test4(self):
        self.assertEqual(list(sideneigh(5, 5, 0, 0)),
                         [(1, 0), (0, 1)])
        self.assertEqual(list(sideneigh(5, 5, 4, 4)),
                         [(3, 4), (4, 3)])
        self.assertEqual(list(sideneigh(5, 5, 0, 4)),
                         [(1, 4), (0, 3)])
        self.assertEqual(list(sideneigh(5, 5, 4, 0)),
                         [(3, 0), (4, 1)])
        self.assertEqual(list(sideneigh(5, 5, 2, 2)),
                         [(1, 2), (3, 2), (2, 1), (2, 3)])

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
