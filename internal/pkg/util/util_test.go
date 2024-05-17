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
