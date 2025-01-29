class SetUnion:
    def __init__(self, *items):
        self.items = set()
        self.parents = {}
        self.sizes = {}
        self.count = 0
        for i in items: self.add(item)

    def add(self, item):
        if not self.has(item):
            self.items.add(item)
            self.parents[item] = item
            self.sizes[item] = 1
            self.count += 1

    def has(self, item):
        return item in self.items

    def find(self, item):
        """ Find the root of a given item. """
        p = self.parents[item]
        if p == item:   # Reached the root
            return p
        return self.find(p)

    def union(self, item1, item2):
        """ Merges two item sets. """
        # Identify roots of the two items
        p1, p2 = self.find(item1), self.find(item2)
        if p1 != p2:
            s1,s2 = self.sizes[p1], self.sizes[p2]
            # Make the "smaller" root a child of the "larger" root
            if s1 >= s2:
                self.sizes[p1] = s1 + s2
                self.parents[p2] = p1
            else:
                self.sizes[p2] = s1 + s2
                self.parents[p1] = p2

    def same_comp(self, item1, item2):
        return self.find(item1) == self.find(item2)

    def byparent(self):
        out = {}
        for item in self.parents.keys():
            parent = self.find(item)
            if parent not in out: out[parent] = []
            out[parent].append(item)
        return out
