package ds

import "log"

type TrieNode struct {
	Ch       rune
	Terminal bool
	Parent   *TrieNode
	Count    int
	Children TrieNodeMap
}

type TrieNodeMap = map[rune]*TrieNode

func NewTrieNode(ch rune, terminal bool, parent *TrieNode) *TrieNode {
	out := TrieNode{
		Ch:       ch,
		Terminal: terminal,
		Parent:   parent,
		Count:    0,
		Children: make(TrieNodeMap),
	}
	return &out
}

// Adds a string (represented as runes) from this node and returns the leaf node of the
// bottom most trie node correpsonding to the last char in the string.
// The terminal flag must be set manually by the caller if needed.
func (t *TrieNode) AddString(str string, offset int) *TrieNode {
	return t.Add([]rune(str), offset)
}
func (t *TrieNode) Add(str []rune, offset int) *TrieNode {
	curr := t
	var child *TrieNode
	for off := offset; off < len(str); off += 1 {
		curr.Count += 1
		ch := str[off]
		child = curr.Children[ch]
		if child == nil {
			child = NewTrieNode(ch, false, curr)
			curr.Children[ch] = child
		}
		curr = child
	}
	if child != nil {
		child.Count += 1
	}
	return curr
}

// Finds the leaf Trienode that corresponds to the last item in the string.
// Usually used to work backwards and other checks.
func (t *TrieNode) FindLeaf(str []rune, offset int) *TrieNode {
	curr := t
	for off := offset; off < len(str); off += 1 {
		if curr.Parent != nil && curr.Count <= 0 {
			panic("0 count nodes must be deleted for non root nodes")
		}
		ch := str[off]
		child := curr.Children[ch]
		if child == nil {
			return nil
		}
		curr = child
	}
	return curr
}

func (t *TrieNode) RemoveString(str string, offset int) bool {
	return t.Remove([]rune(str), offset)
}
func (t *TrieNode) Remove(str []rune, offset int) bool {
	leaf := t.FindLeaf(str, offset)
	if leaf != nil {
		leaf.deccount()
	}
	return leaf != nil
}

func (t *TrieNode) WordsSoFar(reducer func(a []rune, b rune) []rune) []rune {
	if reducer == nil {
		// reduce = lambda a,b: a+b
		reducer = func(a []rune, b rune) []rune { return append(a, b) }
	}
	if t.Parent == nil {
		return []rune{}
	}
	return reducer(t.Parent.WordsSoFar(reducer), t.Ch)
}

// Reduces count of a node and if the count reaches 0 removes itself
// from the parent's child list.
// Recursively calls the parent's counter to be decreased.
func (t *TrieNode) deccount() {
	t.Count -= 1
	if t.Count <= 0 {
		t.Count = 0
		if t.Parent != nil {
			// Remove from the Parent and reduce its Count by one
			_, ok := t.Parent.Children[t.Ch]
			if ok {
				// delete it - dont leave nils lying around
				delete(t.Parent.Children, t.Ch)
			}
		}
	}
	if t.Parent != nil {
		t.Parent.deccount()
	}
}

func (t *TrieNode) IsValid() bool {
	if t.Parent != nil {
		if t.Parent.Children[t.Ch] != t {
			return false
		}
	}

	// check all children
	for _, child := range t.Children {
		if child != nil && !child.IsValid() {
			return false
		}
	}
	return true
}

// Compares two trie nodes to check if they are isomorphic.
func (t *TrieNode) IsIsomorphic(another *TrieNode) bool {
	if t == nil && another == nil {
		return true
	}
	if t == nil || another == nil || len(t.Children) != len(another.Children) ||
		t.Count != another.Count || t.Ch != another.Ch || t.Terminal != another.Terminal {
		log.Println("Failed here: ", t, another)
		return false
	}

	// Compare children
	for key, child := range t.Children {
		a2 := another.Children[key]
		if !child.IsIsomorphic(a2) {
			return false
		}
	}
	return true
}
