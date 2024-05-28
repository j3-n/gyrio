package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testData := []struct {
		Name   string
		Input  []int
		Result int
	}{
		{
			Name:   "Normal",
			Input:  []int{1, 2, 3, 4},
			Result: 10,
		},
		{
			Name:   "SecondNormal",
			Input:  []int{5, 3, 2, 5},
			Result: 15,
		},
		{
			Name:   "Empty",
			Input:  []int{},
			Result: 0,
		},
		{
			Name:   "Negative",
			Input:  []int{10, -10, -15},
			Result: -15,
		},
		{
			Name:   "SecondNegative",
			Input:  []int{-10, -10, 25},
			Result: 5,
		},
	}

	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			sum := Sum(test.Input)
			assert.Equal(t, test.Result, sum, "Test that sum returned the expected result")
		})
	}
}

func TestMod(t *testing.T) {
	testCases := []struct {
		Name   string
		A      int
		B      int
		Result int
	}{
		{
			Name:   "Normal",
			A:      5,
			B:      10,
			Result: 5,
		},
		{
			Name:   "Boundary",
			A:      5,
			B:      5,
			Result: 0,
		},
		{
			Name:   "Zero",
			A:      0,
			B:      5,
			Result: 0,
		},
		{
			Name:   "Negative",
			A:      -1,
			B:      3,
			Result: 2,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Result, Mod(test.A, test.B), "Test that mod returns the expected result")
		})
	}
}

func TestWrapString(t *testing.T) {
	testCases := []struct {
		Name   string
		Input  string
		Width  int
		Output string
	}{
		{
			Name:   "Normal",
			Input:  "abcdefghi",
			Width:  3,
			Output: "abc\ndef\nghi",
		},
		{
			Name:   "SecondNormal",
			Input:  "abcdefghi",
			Width:  4,
			Output: "abcd\nefgh\ni",
		},
		{
			Name:   "1Width",
			Input:  "abc",
			Width:  1,
			Output: "a\nb\nc",
		},
		{
			Name:   "ZeroWidth",
			Input:  "abcdefghi",
			Width:  0,
			Output: "",
		},
		{
			Name:   "NegativeWidth",
			Input:  "abcdefghi",
			Width:  -1,
			Output: "",
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			result := WrapString(test.Input, test.Width)
			assert.Equal(t, test.Output, result, "Test that WrapString returns the expected result")
		})
	}
}
