package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem 6 - https://leetcode.com/problems/zigzag-conversion/
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
// P   A   H   N
// A P L S I I G
// Y   I   R

// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
// string convert(string s, int numRows);
//
// Example 1:
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
//
// Example 2:
//
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
//
// Example 3:
//
// Input: s = "A", numRows = 1
// Output: "A"

// Ideas:
//
// Bruteforce is to simply write this row wise we would to a 2D buffer and then transpose the buffer
// This would need O(n) of space and an extra "transpose" operation
//
// numRows == 1 is a special case and the input is same as the output
//
// The above pattern "repeats".  eg when numRows == 3, PAYP first goes down and then zigs diagonally,
// then the pattern repeats with ALIS, HIRI and NG
// Similary with numRows == 4, we have an extra row + an extra column - resulting in substrings each of 6, eg:
// PAYPAL ISHIRI NG
//
// Let us call this substring length S and can be expressed as:
//
// S = (R - 1) * 2
//
// eg, 2 => 2, 3 => 4, 4 => 6, 5 => 8 and so on
//
// Further more numRows of R can be expressed in terms of numRows of R - 1
// So at row == 0 -
func problem6(s string, numRows int) string {
	buff := bytes.NewBufferString("")
	type Node struct {
		R      int
		Start  int
		End    int
		DebugS string // just for debug
	}

	queue := []Node{
		{Start: 0, End: len(s) - 1, R: numRows, DebugS: s},
	}
	for len(queue) > 0 {
		var newqueue []Node

		// These all should be for the same row
		for _, node := range queue {
			if node.R == 1 {
				buff.WriteString(s[node.Start : node.End+1])
			} else {
				S := (node.R - 1) * 2
				for si := node.Start; si <= node.End; si += S {
					if node.Start <= node.End {
						buff.WriteByte(s[si])
						// handle children
						ns := si + 1
						ne := min(node.End, si+S-1)
						if ns <= ne {
							newqueue = append(newqueue, Node{
								Start:  ns,
								End:    ne,
								R:      node.R - 1,
								DebugS: s[ns : ne+1],
							})
						}
					}
				}
			}
		}
		queue = newqueue
	}
	return buff.String()
}

func TestProblem6(t *testing.T) {
	assert.Equal(t, problem6("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR")
	assert.Equal(t, problem6("PAYPALISHIRING", 4), "PINALSIGYAHRPI")
	assert.Equal(t, problem6("NG", 2), "NG")
}
