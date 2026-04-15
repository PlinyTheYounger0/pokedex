package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello   world     ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		result := cleanInput(c.input)

		if len(result) != len(c.expected) {
			t.Errorf("Expected: %v, Result: %v", c.expected, result)
		}

		for i := range result {
			word := result[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected: %v Result: %v @ %d index", expectedWord, word, i)
			}
		}
	}
}
