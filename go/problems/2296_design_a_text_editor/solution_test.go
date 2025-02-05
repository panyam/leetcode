package main

import (
	"fmt"
	"log"
	"testing"

	_ "embed"

	"github.com/panyam/leetcode/go/utils"
	"github.com/stretchr/testify/assert"
)

// BeginProblemTests

//go:embed "testcases/testcases.full"
var smallTestCases string

//go:embed "testcases/largecase1/commands"
var largeCommands string

//go:embed "testcases/largecase1/args"
var largeArgs string

//go:embed "testcases/largecase1/expected"
var largeExpVals string

type TestCase utils.CommandTestCase

func (tc *TestCase) Run(t *testing.T, id string) {
	log.Println("Test Case: ", id)
	var ed TextEditor
	for i, cmd := range tc.Commands {
		args := tc.Args[i].([]any)
		expval := tc.Expected[i]
		if cmd == "TextEditor" {
			ed = Constructor()
		} else if cmd == "addText" {
			ed.AddText(args[0].(string))
		} else if cmd == "deleteText" {
			assert.Equal(t, int(expval.(float64)), ed.DeleteText(int(args[0].(float64))))
		} else if cmd == "cursorRight" {
			assert.Equal(t, expval, ed.CursorRight(int(args[0].(float64))))
		} else if cmd == "cursorLeft" {
			assert.Equal(t, expval, ed.CursorLeft(int(args[0].(float64))))
		} else {
			log.Fatalf("Invalid command: %s", cmd)
		}
	}
}

func Test2296_MultipleCases(t *testing.T) {
	cases := utils.LoadCases[TestCase]([]byte(smallTestCases))

	for i, tc := range cases {
		tc.Run(t, fmt.Sprintf("smallcase_%d", i))
	}
}

func Test2296_LargeCase(t *testing.T) {
	tc := utils.LoadLargeCase[TestCase](map[string][]byte{"commands": []byte(largeCommands), "args": []byte(largeArgs), "expected": []byte(largeExpVals)})
	tc.Run(t, "largecase")
}

func Test2296_1(t *testing.T) {
	ed := Constructor()
	ed.AddText("leetcode")
	assert.Equal(t, 4, ed.DeleteText(4))
	ed.AddText("practice")
	assert.Equal(t, "etpractice", ed.CursorRight(3))
	assert.Equal(t, "leet", ed.CursorLeft(8))
	assert.Equal(t, 4, ed.DeleteText(10))
	assert.Equal(t, "", ed.CursorLeft(2))
	assert.Equal(t, "practi", ed.CursorRight(6))
}

func Test2296_2(t *testing.T) {
	ed := Constructor()
	ed.AddText("cyberworks")
	assert.Equal(t, 10, ed.DeleteText(21))
	assert.Equal(t, "", ed.CursorLeft(2))
	assert.Equal(t, "", ed.CursorRight(6))
	ed.AddText("aaaaa")
	assert.Equal(t, "aaaa", ed.CursorLeft(1))
}

func Test2296_3(t *testing.T) {
	ed := Constructor()
	assert.Equal(t, "", ed.CursorLeft(1))
	assert.Equal(t, "", ed.CursorRight(4))
	assert.Equal(t, 0, ed.DeleteText(3))
}

func Test2296_4(t *testing.T) {
	ed := Constructor()
	ed.AddText("bxyackuncqzcqo")
	assert.Equal(t, "bx", ed.CursorLeft(12))
	assert.Equal(t, 2, ed.DeleteText(3))
	assert.Equal(t, "", ed.CursorLeft(5))
	ed.AddText("osdhyvqxf")
	assert.Equal(t, "yackuncqzc", ed.CursorRight(10))
}

// EndProblemTests
