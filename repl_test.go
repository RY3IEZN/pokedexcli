package main

import (
	"testing"
)

func TestClearInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello bro",
			expected: []string{
				"hello", "bro",
			},
		},
		{
			input: "hello BRO",
			expected: []string{
				"hello", "bro",
			},
		},
	}

	for _, cs := range cases {
		actual := clearInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("the lengths are not equal: %v vs %v",
				len(actual),
				len(cs.expected),
			)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal to %v",
					actualWord,
					expectedWord)
			}
		}
	}
}
