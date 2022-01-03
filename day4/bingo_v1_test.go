package main

import (
	"testing"
)

func TestBingoBoard_isComplete(t *testing.T) {
	for _, input := range []struct {
		b        bingoBoard
		expected bool
	}{

		{
			b: bingoBoard{
				{{marked: true}},
			},
			expected: false,
		},
		{
			b: bingoBoard{
				{{marked: true}, {marked: true}, {marked: true}, {marked: true}, {marked: true}},
			},
			expected: true,
		},
		{
			b: bingoBoard{
				{{marked: true}},
				{{marked: true}},
				{{marked: true}},
				{{marked: true}},
				{{marked: true}},
			},
			expected: true,
		},
	} {
		t.Run("BingoP1", func(t *testing.T) {
			got := input.b.isComplete()
			if got != input.expected {
				t.Errorf("Expected: %v Got: %v", input.expected, got)
			}
		})
	}
}
