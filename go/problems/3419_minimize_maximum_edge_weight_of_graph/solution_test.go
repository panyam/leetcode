package main

import (
	"fmt"
	"testing"

	_ "embed"

	"github.com/panyam/leetcode/go/utils"
	"github.com/stretchr/testify/assert"
)

// BeginProblemTests

type TestCase struct {
	N         int
	Edges     [][]int
	Threshold int
	Expected  int
}

func (tc *TestCase) Run(t *testing.T, id string) {
	assert.Equal(t, tc.Expected, minMaxWeight(tc.N, tc.Edges, tc.Threshold))
}

//go:embed "testcases.full"
var smallTestCases string

func _Test_MultipleCases(t *testing.T) {
	cases := utils.LoadCases[TestCase]([]byte(smallTestCases))

	for i, tc := range cases {
		tc.Run(t, fmt.Sprintf("smallcase_%d", i))
	}
}

func _Test_LargeCases(t *testing.T) {
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
