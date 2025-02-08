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
	Input    []int `json:"input"`
	Expected bool  `json:"expected"`
}

func (tc *TestCase) Run(t *testing.T, id string) {
	assert.Equal(t, tc.Expected, canJump(tc.Input))
}

//go:embed "testcases.full"
var smallTestCases string

func Test55_MultipleCases(t *testing.T) {
	cases := utils.LoadCases[TestCase]([]byte(smallTestCases))

	for i, tc := range cases {
		tc.Run(t, fmt.Sprintf("smallcase_%d", i))
	}
}

//go:embed "large1.true"
var large1True string

func Test55_LargeCase(t *testing.T) {
	tc := utils.LoadLargeCase[TestCase](map[string][]byte{"input": []byte(large1True)})
	tc.Expected = true
	tc.Run(t, "large1.true")
}
