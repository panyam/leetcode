package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BeginProblemTests

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
