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

//go:embed "testcases.full"
var smallTestCases string

type TestCase utils.CommandTestCase

func (tc *TestCase) Run(t *testing.T, id string) {
	log.Println("Test Case: ", id)
	var ms MaxStack
	for i, cmd := range tc.Commands {
		args := tc.Args[i].([]any)
		expval := tc.Expected[i]
		log.Println("Running command: ", cmd)
		if cmd == "MaxStack" {
			ms = Constructor()
		} else if cmd == "push" {
			ms.Push(int(args[0].(float64)))
		} else if cmd == "pop" {
			assert.Equal(t, int(expval.(float64)), ms.Pop())
		} else if cmd == "top" {
			assert.Equal(t, int(expval.(float64)), ms.Top())
		} else if cmd == "peekMax" {
			assert.Equal(t, int(expval.(float64)), ms.PeekMax())
		} else if cmd == "popMax" {
			assert.Equal(t, int(expval.(float64)), ms.PopMax())
		} else {
			log.Fatalf("Invalid command: %s", cmd)
		}
	}
}

func _Test716_MultipleCases(t *testing.T) {
	cases := utils.LoadCases[TestCase]([]byte(smallTestCases))

	for i, tc := range cases {
		tc.Run(t, fmt.Sprintf("smallcase_%d", i))
	}
}

/*
//go:embed "testcases/largecase1/commands"
var largeCommands string

//go:embed "testcases/largecase1/args"
var largeArgs string

//go:embed "testcases/largecase1/expected"
var largeExpVals string

func Test716_LargeCase(t *testing.T) {
	tc := utils.LoadLargeCase[TestCase](map[string][]byte{"commands": []byte(largeCommands), "args": []byte(largeArgs), "expected": []byte(largeExpVals)})
	tc.Run(t, "largecase")
}
*/
