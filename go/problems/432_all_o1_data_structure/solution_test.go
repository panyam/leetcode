package main

import (
	"fmt"
	"log"
	"testing"

	_ "embed"

	"github.com/panyam/leetcode/go/utils"
)

// BeginProblemTests

//go:embed "testcases.full"
var smallTestCases string

type TestCase utils.CommandTestCase

func (tc *TestCase) Run(t *testing.T, id string) {
	log.Println("Test Case: ", id)
	var it AllOne
	for i, cmd := range tc.Commands {
		args := tc.Args[i].([]any)
		expval := tc.Expected[i]
		// log.Println("Running command: ", cmd)
		if cmd == "AllOne" {
			it = Constructor()
		} else if cmd == "inc" {
			it.Inc(args[0].(string))
		} else if cmd == "dec" {
			it.Dec(args[0].(string))
		} else if cmd == "getMaxKey" {
			// it.Print("getmax: ")
			res := it.GetMaxKey()
			if !utils.ExpectValOrContains(t, expval, res) {
				it.Print(fmt.Sprintf("Mismatch in command: %d, %s", i, cmd))
			}
		} else if cmd == "getMinKey" {
			res := it.GetMinKey()
			if !utils.ExpectValOrContains(t, expval, res) {
				log.Printf("Mismatch in command: %d, %s", i, cmd)
				it.Print("")
			}
		} else {
			log.Fatalf("Invalid command: %s", cmd)
		}
	}
}

func Test716_MultipleCases(t *testing.T) {
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
