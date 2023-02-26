
from functools import total_ordering

@total_ordering

class Comparable:
    """ A way to wrap an arbitrary object with our comparable function. """
    def __init__(self, val, ltfunc, eqfunc):
        self.val = val
        self.ltfunc = ltfunc
        self.eqfunc = eqfunc

    def __lt__(self, another):
        return self.ltfunc(self.val, another.val)

    def __eq__(self, another):
        return self.eqfunc(self.val, another.val)

    def __hash__(self):
        return hash(self.val)
