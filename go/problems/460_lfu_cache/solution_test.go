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
	var it LFUCache
	for i, cmd := range tc.Commands {
		args := tc.Args[i].([]any)
		expval := tc.Expected[i]
		log.Println("Running command: ", cmd, args)
		if cmd == "LFUCache" {
			it = Constructor(int(args[0].(float64)))
			it.Debug("test")
		} else if cmd == "put" {
			it.Put(int(args[0].(float64)), int(args[1].(float64)))
		} else if cmd == "get" {
			assert.Equal(t, int(expval.(float64)), it.Get(int(args[0].(float64))))
		} else {
			log.Fatalf("Invalid command: %s", cmd)
		}
		it.Debug("xxxxxxxx")
	}
}

func Test146_MultipleCases(t *testing.T) {
	cases := utils.LoadCases[TestCase]([]byte(smallTestCases))

	for i, tc := range cases {
		tc.Run(t, fmt.Sprintf("smallcase_%d", i))
		return
	}
}

//go:embed "testcases/large1.commands"
var commandsLargeCase1 string

//go:embed "testcases/large1.args"
var argsLargeCase1 string

//go:embed "testcases/large1.expected"
var expValsLargeCase1 string

func Test146_LargeCase(t *testing.T) {
	tc := utils.LoadLargeCase[TestCase](map[string][]byte{"commands": []byte(commandsLargeCase1), "args": []byte(argsLargeCase1), "expected": []byte(expValsLargeCase1)})
	tc.Run(t, "largecase")
}
