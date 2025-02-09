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

type TestCase struct {
	Maze     [][]int
	Ball     []int
	Hole     []int
	Expected string
}

func (tc *TestCase) Run(t *testing.T, id string) bool {
	return assert.Equal(t, tc.Expected, findShortestWay(tc.Maze, tc.Ball, tc.Hole))
}

//go:embed "testcases.full"
var smallTestCases string

func Test_MultipleCases(t *testing.T) {
	cases := utils.LoadCases[TestCase]([]byte(smallTestCases))

	for i, tc := range cases {
		id := fmt.Sprintf("smallcase_%d", i)
		if !tc.Run(t, id) {
			log.Println("Case Failed: ", id)
		}
	}
}

func Test_LargeCases(t *testing.T) {
	// Find a way to load and run large files
}

/*
//go:embed "large1.true"
var large1True string

func Test_LargeCase(t *testing.T) {
	tc := utils.LoadLargeCase[TestCase](map[string][]byte{"input": []byte(large1True)})
	tc.Expected = 10
	tc.Run(t, "large1.true")
}
*/
