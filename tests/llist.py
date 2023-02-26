from algos.llist import *

import unittest
class TestMethods(unittest.TestCase):
    def test_listlen(self):
        self.assertEqual(listlen(None), 0)

        h,t = llmake() ; listlen(h)
        self.assertEqual(listlen(h), 0)

        h,t = llmake(1,2,3,4,5) ; listlen(h)
        self.assertEqual(listlen(h), 5)

    def test_listcmp(self):
        h1,t1 = llmake(1,2,3,4,5)
        h2,t2 = llmake(1,2,3,4,5)
        self.assertEqual(listcmp(h1, h2), 0)

        h1,t1 = llmake(1,2,3,4,5)
        h2,t2 = llmake(1,2,3,4,5,6)
        self.assertEqual(listcmp(h1, h2), -1)

        h1,t1 = llmake(1,2,3,4,5)
        h2,t2 = llmake(1,2,3,4)
        self.assertEqual(listcmp(h1, h2), 1)

        h1,t1 = llmake(1,2,3,4,5)
        h2,t2 = llmake(1,2,3,4,7)
        listcmp(h1, h2)
        self.assertEqual(listcmp(h1, h2), -2)

    def test_lltail(self):
        h1,t1 = llmake(1,2,3,4,5)
        self.assertEqual(lltail(h1).val, 5)

    def test_llmin(self):
        h1,t1 = llmake(10, 4, 9, 11, 1, 200)
        self.assertEqual(llmin(h1).val, 1)
    
    def test_llmid(self):
        h1,t1 = llmake(1)
        prev,mid = llmid(h1)
        self.assertEqual((prev, mid.val), (None, 1))

        h1,t1 = llmake(1, 2)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (1, 2))

        h1,t1 = llmake(10, 4, 9, 11, 1, 200)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (9, 11))

        h1,t1 = llmake(10, 4, 9, 11, 1, 200, 300)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (9, 11))

    def test_llmid_circular(self):
        h1,t1 = llcircmake(1)
        prev,mid = llmid(h1)
        self.assertEqual((prev, mid.val), (None, 1))

        h1,t1 = llcircmake(1, 2)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (1, 2))

        h1,t1 = llcircmake(10, 4, 9, 11, 1, 200)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (9, 11))

        h1,t1 = llcircmake(10, 4, 9, 11, 1, 200, 300)
        prev,mid = llmid(h1)
        self.assertEqual((prev.val, mid.val), (9, 11))

    def test_lldel(self):
        # Insert into empty list
        h1, t1 = llmake(1,2,3,4,5)

        # remove head
        h2, t2 = lldel(h1, t1, h1)
        self.assertEqual(llvalues(h2), [2, 3, 4, 5])
        self.assertEqual(h2.val, 2)
        self.assertEqual(t2.val, 5)

        # remove tail
        h3, t3 = lldel(h2, t2, t2, doubly=True)
        self.assertEqual(llvalues(h3), [2, 3, 4])
        self.assertEqual(h3.val, 2)
        self.assertEqual(t3.val, 4)

        # val from the middle
        h4, t4 = lldel(h3, t3, h3.next, doubly=True)
        self.assertEqual(llvalues(h3), [2, 4])
        self.assertEqual(h4.val, 2)
        self.assertEqual(t4.val, 4)

    def test_llinsert(self):
        # Insert into empty list
        n1 = Node(10)
        h1, t1 = llinsert(None, None, n1, doubly=True)
        self.assertEqual(h1, n1)
        self.assertEqual(t1, n1)
        self.assertEqual(llvalues(h1), [10])

        # Insert before head
        n2 = Node(5)
        nh, nt = llinsert(h1, t1, n2, None, doubly=True)
        self.assertEqual(llvalues(nh), [5, 10])
        self.assertEqual(nh, n2)
        self.assertEqual(nt, t1)

        # Insert after tail
        n3 = Node(15)
        h3, t3 = llinsert(nh, nt, n3, nt, doubly=True)
        self.assertEqual(llvalues(h3), [5, 10, 15])
        self.assertEqual(h3, nh)
        self.assertEqual(t3, n3)

        # Insert at an arbitrary position
        n4 = Node(12)
        h4, t4 = llinsert(h3, t3, n4, n1, doubly=True)
        self.assertEqual(llvalues(h4), [5, 10, 12, 15])
        self.assertEqual(h4, nh)
        self.assertEqual(t4, n3)

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
