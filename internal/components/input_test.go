package components

import (
	"testing"

	ui "github.com/gizak/termui/v3"
	"github.com/stretchr/testify/assert"
)

type TestInput struct {
	Input  string
	Result string
}

func TestTyping(t *testing.T) {
	testInputs := []TestInput{
		{
			Input:  "H",
			Result: "H",
		},
		{
			Input:  "e",
			Result: "He",
		},
		{
			Input:  "l",
			Result: "Hel",
		},
		{
			Input:  "<Control>",
			Result: "Hel",
		},
		{
			Input:  "Wrong",
			Result: "Hel",
		},
		{
			Input:  "l",
			Result: "Hell",
		},
		{
			Input:  "o",
			Result: "Hello",
		},
	}

	runTestInputSet(t, testInputs)
}

func TestNavigatingText(t *testing.T) {
	testInputs := []TestInput{
		{
			Input:  "<Left>",
			Result: "",
		},
		{
			Input:  "<Right>",
			Result: "",
		},
		{
			Input:  "<Backspace>",
			Result: "",
		},
		{
			Input:  "A",
			Result: "A",
		},
		{
			Input:  "<Left>",
			Result: "A",
		},
		{
			Input:  "<Left>",
			Result: "A",
		},
		{
			Input:  "B",
			Result: "BA",
		},
		{
			Input:  "<Space>",
			Result: "B A",
		},
		{
			Input:  "<Right>",
			Result: "B A",
		},
		{
			Input:  "<Right>",
			Result: "B A",
		},
		{
			Input:  "C",
			Result: "B AC",
		},
		{
			Input:  "<Backspace>",
			Result: "B A",
		},
		{
			Input:  "<Left>",
			Result: "B A",
		},
		{
			Input:  "<Backspace>",
			Result: "BA",
		},
		{
			Input:  "<Left>",
			Result: "BA",
		},
		{
			Input:  "<Backspace>",
			Result: "BA",
		},
	}

	runTestInputSet(t, testInputs)
}

func runTestInputSet(t *testing.T, testInputs []TestInput) {
	input := NewInput()
	for _, test := range testInputs {
		t.Run("TestInputs", func(t *testing.T) {
			input.KeyboardEvent(&ui.Event{
				Type: ui.KeyboardEvent,
				ID:   test.Input,
			})
			assert.Equal(t, test.Result, input.CurrentText, "Test that input contains the expected text")
		})
	}
}
