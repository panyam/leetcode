
class TrieNode:
    def __init__(self, char, terminal=False, parent=None):
        self.char = char 
        self.terminal = terminal
        self.parent = parent
        self.count = 0
        self.children = {}


    def add_string(self, string, offset=0):
        """ Adds a string from this node and returns the leaf node of the
        bottom most trie node correpsonding to the last char in the string.
        The terminal flag must be set manually by the caller if needed.
        """
        curr = self
        child = None
        for off in range(offset, len(string)):
            curr.count += 1
            char = string[off]
            child = curr.children.get(char, None)
            if not child:
                child = curr.children[char] = TrieNode(char, False, curr)
            curr = child
        if child: child.count += 1
        return curr

    def find_leaf(self, string, offset=0):
        """ Finds the leaf Trienode that corresponds to the last item in
        the string.
        Usually used to work backwards and other checks.
        """
        curr = self
        for off in range(offset, len(string)):
            assert curr.parent is None or curr.count > 0, "0 count nodes must be deleted for non root nodes"
            char = string[off]
            child = curr.children.get(char, None)
            if not child:
                return None
            curr = child
        return curr

    def remove_string(self, string, offset=0):
        leaf = self.find_leaf(string, offset)
        if leaf: leaf._deccount()
        return leaf is not None

    def wordsofar(self, reduce=None):
        if not reduce:
            reduce = lambda a,b: a+b
        if self.parent is None:
            return self.char
        return reduce(self.parent.wordsofar(reduce), self.char)

    def _deccount(self):
        """ Reduces count of a node and if the count reaches 0 removes itself
        from the parent's child list.
        Recursively calls the parent's counter to be decreased.
        """
        self.count -= 1
        if self.count <= 0:
            self.count = 0
            if self.parent:
                # Remove from the parent and reduce its count by one
                del self.parent.children[self.char]
        if self.parent:
            self.parent._deccount()

