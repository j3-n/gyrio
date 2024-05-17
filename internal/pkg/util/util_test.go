package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		Title  string
		Input  []int
		Result int
	}{
		{
			Title:  "Normal",
			Input:  []int{1, 2, 3, 4},
			Result: 10,
		},
		{
			Title:  "SecondNormal",
			Input:  []int{5, 3, 2, 5},
			Result: 15,
		},
		{
			Title:  "Empty",
			Input:  []int{},
			Result: 0,
		},
		{
			Title:  "Negative",
			Input:  []int{10, -10, -15},
			Result: -15,
		},
		{
			Title:  "SecondNegative",
			Input:  []int{-10, -10, 25},
			Result: 5,
		},
	}

	for _, test := range testCases {
		t.Run(test.Title, func(t *testing.T) {
			sum := Sum(test.Input)
			assert.Equal(t, test.Result, sum, "Test that sum returned the expected result")
		})
	}
}
