package components

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumnWidths(t *testing.T) {
	testData := []struct {
		Name       string
		Titles     []string
		Data       [][]string
		Result     []int
		ShouldFail bool
	}{
		{
			Name:   "NormalOneRow",
			Titles: []string{"a", "b", "c", "d"},
			Data: [][]string{
				{"a", "bb", "ccc", "dddd"},
			},
			Result:     []int{1, 2, 3, 4},
			ShouldFail: false,
		},
		{
			Name:   "SecondNormalOneRow",
			Titles: []string{"a", "b", "c", "d"},
			Data: [][]string{
				{"aaaa", "bbb", "ccc", "dddd"},
			},
			Result:     []int{4, 3, 3, 4},
			ShouldFail: false,
		},
		{
			Name:   "NormalOneEmptyRow",
			Titles: []string{"a", "b", "c", "d"},
			Data: [][]string{
				{"", "", "", ""},
			},
			Result:     []int{1, 1, 1, 1},
			ShouldFail: false,
		},
		{
			Name:   "NormalMultipleRows",
			Titles: []string{"a", "b", "c", "d"},
			Data: [][]string{
				{"a", "bb", "c", "d"},
				{"a", "bb", "ccc", "dddd"},
				{"aa", "b", "c", "dd"},
			},
			Result:     []int{2, 2, 3, 4},
			ShouldFail: false,
		},
		{
			Name:   "NormalLargeTitle",
			Titles: []string{"aaaaa", "b", "c", "d"},
			Data: [][]string{
				{"a", "bb", "c", "d"},
				{"a", "bb", "ccc", "dddd"},
				{"aa", "b", "c", "dd"},
			},
			Result:     []int{5, 2, 3, 4},
			ShouldFail: false,
		},
		{
			Name:   "ColumnMismatchTooFew",
			Titles: []string{"a", "b", "c", "d"},
			Data: [][]string{
				{"a", "bb", "c", "d"},
				{"a", "bb", "ccc"},
				{"aa", "b", "c", "dd"},
			},
			ShouldFail: true,
		},
		{
			Name:   "ColumnMismatchTooMany",
			Titles: []string{"a", "b", "c", "d"},
			Data: [][]string{
				{"a", "bb", "c", "d"},
				{"a", "bb", "ccc", "dd"},
				{"aa", "b", "c", "dd", "ee"},
			},
			ShouldFail: true,
		},
		{
			Name:       "NoData",
			Titles:     []string{"a", "b", "c", "d"},
			Data:       [][]string{},
			Result:     []int{1, 1, 1, 1},
			ShouldFail: false,
		},
		{
			Name:   "NoColumnTitles",
			Titles: []string{},
			Data: [][]string{
				{"a", "b", "c", "d"},
			},
			ShouldFail: true,
		},
	}

	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			r, err := computeWidths(test.Titles, test.Data)
			if test.ShouldFail {
				assert.Error(t, err, "Test that computeWidths fails when expected")
			} else {
				assert.True(t, reflect.DeepEqual(r, test.Result), "Test that computeWidths returns the expected result")
			}
		})
	}
}
