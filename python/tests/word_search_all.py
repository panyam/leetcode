
from algos.grid import sideneigh
from algos.bfs import bfs
from algos.dfs import DFS, simpledfs
from algos.trie import TrieNode
from typing import List
import unittest
import ipdb

def solution(board: List[List[str]], words: List[str]) -> List[str]:
    if not board or not board[0]: return []
    R = len(board)
    C = len(board[0])

    strie = {}
    for w in words:
        if w[0] not in strie:
            strie[w[0]] = []
        strie[w[0]].append(w)

    trie = TrieNode("")
    for w in words:
        leaf = trie.add_string(w)
        leaf.terminal = True

    out = set()

    def dodfs2(r,c,node):
        def neigh(loc, trienode, parents):
            r,c = loc
            for nr,nc in sideneigh(R, C, r, c):
                # Only yield if part of the trie-node
                ch = board[nr][nc]
                if ch in trienode.children:
                    yield (nr,nc), trienode.children[ch]
        for (nr,nc),nnode in simpledfs(neigh, (r,c), node):
            if nnode.terminal:
                out.add(nnode.wordsofar())

    def dodfs3(r,c,word):
        if word in out: return

        L = len(word)
        def neigh2(loc, wi, parents):
            r,c = loc
            if wi < L - 1:
                for nr,nc in sideneigh(R, C, r, c):
                    # Only yield if part of the trie-node
                    ch = board[nr][nc]
                    if ch == word[wi + 1]:
                        yield (nr,nc), wi + 1

        for (nr,nc),wi in simpledfs(neigh2, (r,c), 0):
            if wi >= L - 1:
                out.add(word)


    return
    if True:
        for ch,tnode in trie.children.items():
            for r in range(R):
                for c in range(C):
                    if board[r][c] == ch:
                        dodfs2(r,c,tnode)
    else:
        for r in range(R):
            for c in range(C):
                ch = board[r][c]
                if False:
                    if ch in trie.children:
                        tnode = trie.children[ch]
                        print("Matching: ", r,c,ch, tnode.key, tnode.children)
                        dodfs2(r, c, tnode)
                else:
                    if ch in strie:
                        for word in strie[ch]:
                            dodfs3(r,c,word)
    return list(out)

class TestMethods(unittest.TestCase):
    def test_cases(self):
        # self.assertEqual(solve_surrounded_regions(grid),solution)
        for case, expected in cases:
            found = solution(*case)
            self.assertEqual(list(sorted(found)), list(sorted(expected)))

cases = [
    (
        (
            [["a","a"]],
            ["aaa"]
        ),
        []
    ),
    (
        (
            [["a"]],
            ["a"]
        ),
        ["a"]
    ),
    (
        (
            [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
            ["oath","pea","eat","rain"]
        ),
        ["oath","eat"]
    ),
    (
        (
            [["a","b"],["c","d"]],
            ["abcb"]
        ),
        []
    ),
    (
        (
            [["o","a","a","n"],
             ["e","t","a","e"],
             ["i","h","k","r"],
             ["i","f","l","v"]],
            ["oath","pea","eat","rain","hklf", "hf"]
        ),
        ["oath","eat","hklf","hf"]
    ),
    (
        (
            [["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"],["a","a","a","a","a","a","a","a","a","a","a","a"]],
            ["lllllll","fffffff","ssss","s","rr","xxxx","ttt","eee","ppppppp","iiiiiiiii","xxxxxxxxxx","pppppp","xxxxxx","yy","jj","ccc","zzz","ffffffff","r","mmmmmmmmm","tttttttt","mm","ttttt","qqqqqqqqqq","z","aaaaaaaa","nnnnnnnnn","v","g","ddddddd","eeeeeeeee","aaaaaaa","ee","n","kkkkkkkkk","ff","qq","vvvvv","kkkk","e","nnn","ooo","kkkkk","o","ooooooo","jjj","lll","ssssssss","mmmm","qqqqq","gggggg","rrrrrrrrrr","iiii","bbbbbbbbb","aaaaaa","hhhh","qqq","zzzzzzzzz","xxxxxxxxx","ww","iiiiiii","pp","vvvvvvvvvv","eeeee","nnnnnnn","nnnnnn","nn","nnnnnnnn","wwwwwwww","vvvvvvvv","fffffffff","aaa","p","ddd","ppppppppp","fffff","aaaaaaaaa","oooooooo","jjjj","xxx","zz","hhhhh","uuuuu","f","ddddddddd","zzzzzz","cccccc","kkkkkk","bbbbbbbb","hhhhhhhhhh","uuuuuuu","cccccccccc","jjjjj","gg","ppp","ccccccccc","rrrrrr","c","cccccccc","yyyyy","uuuu","jjjjjjjj","bb","hhh","l","u","yyyyyy","vvv","mmm","ffffff","eeeeeee","qqqqqqq","zzzzzzzzzz","ggg","zzzzzzz","dddddddddd","jjjjjjj","bbbbb","ttttttt","dddddddd","wwwwwww","vvvvvv","iii","ttttttttt","ggggggg","xx","oooooo","cc","rrrr","qqqq","sssssss","oooo","lllllllll","ii","tttttttttt","uuuuuu","kkkkkkkk","wwwwwwwwww","pppppppppp","uuuuuuuu","yyyyyyy","cccc","ggggg","ddddd","llllllllll","tttt","pppppppp","rrrrrrr","nnnn","x","yyy","iiiiiiiiii","iiiiii","llll","nnnnnnnnnn","aaaaaaaaaa","eeeeeeeeee","m","uuu","rrrrrrrr","h","b","vvvvvvv","ll","vv","mmmmmmm","zzzzz","uu","ccccccc","xxxxxxx","ss","eeeeeeee","llllllll","eeee","y","ppppp","qqqqqq","mmmmmm","gggg","yyyyyyyyy","jjjjjj","rrrrr","a","bbbb","ssssss","sss","ooooo","ffffffffff","kkk","xxxxxxxx","wwwwwwwww","w","iiiiiiii","ffff","dddddd","bbbbbb","uuuuuuuuu","kkkkkkk","gggggggggg","qqqqqqqq","vvvvvvvvv","bbbbbbbbbb","nnnnn","tt","wwww","iiiii","hhhhhhh","zzzzzzzz","ssssssssss","j","fff","bbbbbbb","aaaa","mmmmmmmmmm","jjjjjjjjjj","sssss","yyyyyyyy","hh","q","rrrrrrrrr","mmmmmmmm","wwwww","www","rrr","lllll","uuuuuuuuuu","oo","jjjjjjjjj","dddd","pppp","hhhhhhhhh","kk","gggggggg","xxxxx","vvvv","d","qqqqqqqqq","dd","ggggggggg","t","yyyy","bbb","yyyyyyyyyy","tttttt","ccccc","aa","eeeeee","llllll","kkkkkkkkkk","sssssssss","i","hhhhhh","oooooooooo","wwwwww","ooooooooo","zzzz","k","hhhhhhhh","aaaaa","mmmmm"]
        ),
        [],
    ),
]

if __name__ == '__main__':
    # Run with python -m unittest *.py
    unittest.main()
