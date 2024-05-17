package components

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumnWidths(t *testing.T) {
	testData := []struct {
		Name       string
		Data       [][]string
		Result     []int
		ShouldFail bool
	}{
		{
			Name: "NormalOneRow",
			Data: [][]string{
				{"a", "bb", "ccc", "dddd"},
			},
			Result:     []int{1, 2, 3, 4},
			ShouldFail: false,
		},
		{
			Name: "NormalOneEmptyRow",
			Data: [][]string{
				{"", "", "", ""},
			},
			Result:     []int{0, 0, 0, 0},
			ShouldFail: false,
		},
		{
			Name: "NormalMultipleRows",
			Data: [][]string{
				{"a", "bb", "c", "d"},
				{"a", "bb", "ccc", "dddd"},
				{"aa", "b", "c", "dd"},
			},
			Result:     []int{2, 2, 3, 4},
			ShouldFail: false,
		},
		{
			Name: "ColumnMismatchTooFew",
			Data: [][]string{
				{"a", "bb", "c", "d"},
				{"a", "bb", "ccc"},
				{"aa", "b", "c", "dd"},
			},
			Result:     []int{2, 2, 3, 4},
			ShouldFail: true,
		},
		{
			Name: "ColumnMismatchTooMany",
			Data: [][]string{
				{"a", "bb", "c", "d"},
				{"a", "bb", "ccc", "dd"},
				{"aa", "b", "c", "dd", "ee"},
			},
			Result:     []int{2, 2, 3, 4},
			ShouldFail: true,
		},
	}

	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			r, err := computeWidths(test.Data)
			if test.ShouldFail {
				assert.Error(t, err, "Test that computeWidths fails when expected")
			} else {
				assert.True(t, reflect.DeepEqual(r, test.Result), "Test that computeWidths returns the expected result")
			}
		})
	}
}
