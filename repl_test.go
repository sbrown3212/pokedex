package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"empty": {
			input:    "   ",
			expected: []string{},
		},
		"split_on_whitespace": {
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		"strip_whitespace": {
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		"lowercase": {
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("expected: %v, actual: %v", len(tc.expected), len(actual))
			}
			for i := range actual {
				word := actual[i]
				expectedWord := tc.expected[i]
				if word != expectedWord {
					t.Errorf("Index: %v, expected word: %v, actual word: %v", i, word, expectedWord)
				}
			}
		})
	}
}
