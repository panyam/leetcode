
from tests.utils import run_cases
from algos.unionfind import SetUnion
from typing import List
import unittest
import ipdb

def solution(accounts: List[List[str]]) -> List[List[str]]:
    setu = SetUnion()

    names = {}
    for index,account in enumerate(accounts):
        name, emails = account[0], account[1:]

        for email in emails:
            if email in names:
                assert names[email] == name
            names[email] = name
            setu.add(email)
            setu.union(emails[0], email)

    out = []
    emailsbyparent = setu.byparent()
    for parent, items in emailsbyparent.items():
        items.sort()
        out.append([names[parent]] + items)
    print("Emails: ", out)
    return out

cases = [
    (
        [["John","johnsmith@mail.com","john_newyork@mail.com"],
         ["John","johnsmith@mail.com","john00@mail.com"],
         ["Mary","mary@mail.com"],
         ["John","johnnybravo@mail.com"]],
        [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],
         ["Mary","mary@mail.com"],
         ["John","johnnybravo@mail.com"]]
    ),
    (
        [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"]],
        [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]]
    ),
    (
        [["David","David0@m.co","David1@m.co"],
         ["David","David3@m.co","David4@m.co"],
         ["David","David4@m.co","David5@m.co"],
         ["David","David2@m.co","David3@m.co"],
         ["David","David1@m.co","David2@m.co"]],
        [["David","David0@m.co","David1@m.co","David2@m.co","David3@m.co","David4@m.co","David5@m.co"]]
    ),
]

class _TestMethods(unittest.TestCase):
    def test_it(self):
        # And now in dfs
        run_cases(self, cases, solution, sort=True)

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
