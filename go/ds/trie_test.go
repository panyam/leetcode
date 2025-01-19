package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieInit(t *testing.T) {
	tn := NewTrieNode(rune(0), false, nil)
	assert.Equal(t, tn.Children, TrieNodeMap{})
	assert.Equal(t, tn.Count, 0)
	assert.False(t, tn.Terminal)

	tn.AddString("app", 0).Terminal = true
	tn.AddString("apple", 0).Terminal = true
	tn.AddString("bad", 0).Terminal = true
	tn.AddString("code", 0).Terminal = true
	tn.AddString("cody", 0).Terminal = true
	tn.AddString("cobb", 0).Terminal = true

	expected := &TrieNode{
		Ch:       0,
		Terminal: false,
		Count:    6,
		Children: TrieNodeMap{
			'a': {
				Ch:       'a',
				Terminal: false, Count: 2,
				Children: TrieNodeMap{
					'p': {Terminal: false, Count: 2, Ch: 'p',
						Children: TrieNodeMap{
							'p': {Terminal: true, Count: 2, Ch: 'p',
								Children: TrieNodeMap{
									'l': {Terminal: false, Count: 1, Ch: 'l',
										Children: TrieNodeMap{
											'e': {Terminal: true, Count: 1, Ch: 'e',
												Children: TrieNodeMap{}}}}}}}},
				},
			},
			'b': {Terminal: false, Count: 1, Ch: 'b',
				Children: TrieNodeMap{
					'a': {Terminal: false, Count: 1, Ch: 'a',
						Children: TrieNodeMap{
							'd': {Terminal: true, Count: 1, Ch: 'd',
								Children: TrieNodeMap{}}}}}},
			'c': {Terminal: false, Count: 3, Ch: 'c',
				Children: TrieNodeMap{
					'o': {Terminal: false, Count: 3, Ch: 'o',
						Children: TrieNodeMap{
							'd': {Terminal: false, Count: 2, Ch: 'd',
								Children: TrieNodeMap{
									'e': {Terminal: true, Count: 1, Ch: 'e',
										Children: TrieNodeMap{}},
									'y': {Terminal: true, Count: 1, Ch: 'y',
										Children: TrieNodeMap{}}}},
							'b': {Terminal: false, Ch: 'b',
								Count: 1,
								Children: TrieNodeMap{
									'b': {Terminal: true, Count: 1, Ch: 'b',
										Children: TrieNodeMap{}}}}}}}}},
	}

	assert.True(t, tn.IsIsomorphic(expected))

	// Nothing should change when removing a non-existent val
	assert.False(t, tn.RemoveString("application", 0))
	assert.True(t, tn.IsIsomorphic(expected))
}

func TestTrieRemoval(t *testing.T) {
	tn := NewTrieNode(rune(0), false, nil)
	assert.Equal(t, tn.Children, TrieNodeMap{})
	assert.Equal(t, tn.Count, 0)
	assert.False(t, tn.Terminal)

	tn.AddString("app", 0).Terminal = true
	tn.AddString("apple", 0).Terminal = true
	tn.AddString("bad", 0).Terminal = true
	tn.AddString("code", 0).Terminal = true
	tn.AddString("cody", 0).Terminal = true
	tn.AddString("cobb", 0).Terminal = true

	assert.True(t, tn.RemoveString("app", 0))
	expected := &TrieNode{
		Terminal: false, Count: 5, Ch: 0,
		Children: TrieNodeMap{
			'a': {Terminal: false, Count: 1, Ch: 'a',
				Children: TrieNodeMap{
					'p': {Terminal: false, Count: 1, Ch: 'p',
						Children: TrieNodeMap{
							'p': {Terminal: true, Count: 1, Ch: 'p',
								Children: TrieNodeMap{
									'l': {Terminal: false, Count: 1, Ch: 'l',
										Children: TrieNodeMap{
											'e': {Terminal: true, Count: 1, Ch: 'e',
												Children: TrieNodeMap{}}}}}}}}}},
			'b': {Terminal: false, Count: 1, Ch: 'b',
				Children: TrieNodeMap{
					'a': {Terminal: false, Count: 1, Ch: 'a',
						Children: TrieNodeMap{
							'd': {Terminal: true, Count: 1, Ch: 'd',
								Children: TrieNodeMap{}}}}}},
			'c': {Terminal: false, Count: 3, Ch: 'c',
				Children: TrieNodeMap{
					'o': {Terminal: false, Count: 3, Ch: 'o',
						Children: TrieNodeMap{
							'd': {Terminal: false, Count: 2, Ch: 'd',
								Children: TrieNodeMap{
									'e': {Terminal: true, Count: 1, Ch: 'e',
										Children: TrieNodeMap{}},
									'y': {Terminal: true, Count: 1, Ch: 'y',
										Children: TrieNodeMap{}}}},
							'b': {Terminal: false, Count: 1, Ch: 'b',
								Children: TrieNodeMap{
									'b': {Terminal: true, Count: 1, Ch: 'b',
										Children: TrieNodeMap{}}}}}}}}}}
	assert.True(t, tn.IsIsomorphic(expected))

	assert.Equal(t, tn.FindLeaf([]rune("apple"), 0).WordsSoFar(nil), []rune("apple"))
}

/*
import unittest
class TestMethods(unittest.TestCase):
    def test_init(self):
        tn = TrieNode("")
        self.assertEqual(debuginfo(tn), {'children': {}, 'count': 0, 'terminal': false})

        tn.add_string("app").terminal = true
        tn.add_string("apple").terminal = true
        tn.add_string("bad").terminal = true
        tn.add_string("code").terminal = true
        tn.add_string("cody").terminal = true
        tn.add_string("cobb").terminal = true

        self.assertEqual(tn.remove_string("application"), false)
        # Nothing should change when removing a non-existent val
        self.assertEqual(debuginfo(tn),
                         {'children':
                          {'a': {'children':
                            {'p': {'children':
                              {'p': {'children':
                                {'l': {'children':
                                  {'e': {'children': {}, 'count': 1, 'terminal': true}},
                                  'count': 1, 'terminal': false }},
                                'count': 2, 'terminal': true}},
                              'count': 2, 'terminal': false}},
                            'count': 2, 'terminal': false},
                           'b': {'children':
                                 {'a': {'children':
                                        {'d': {'children': {}, 'count': 1, 'terminal': true}},
                                       'count': 1, 'terminal': false}},
                                 'count': 1, 'terminal': false},
                           'c': {'children':
                                 {'o': {'children':
                                        {'b': {'children':
                                               {'b': {'children': {}, 'count': 1, 'terminal': true}},
                                               'count': 1, 'terminal': false},
                                         'd': {'children':
                                               {'e': {'children': {}, 'count': 1, 'terminal': true},
                                                'y': {'children': {}, 'count': 1, 'terminal': true}},
                                               'count': 2, 'terminal': false}},
                                        'count': 3, 'terminal': false}},
                                 'count': 3, 'terminal': false}},
                          'count': 6, 'terminal': false})

def debuginfo(tnode: TrieNode):
    return {
        "terminal": tnode.terminal,
        "count": tnode.count,
        "children": {
            k: debuginfo(v) for k,v in tnode.children.items()
        }
    }
*/
