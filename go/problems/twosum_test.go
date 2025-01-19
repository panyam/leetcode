package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	assert.Equal(t, twoSum(nums, 9), []int{0, 1})
}
