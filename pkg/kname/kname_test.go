package kname

import (
	"testing"
)

type ShortenTest struct {
	name     string
	attempt  int
	expected string
}

func TestGetShorten(t *testing.T) {
	tests := []ShortenTest{
		{"", 1, ""},
		{"John", 0, ""},
		{"John", 1, "john"},
		{"John", 2, "john_1"},
		{"John", 3, "john_2"},
		{"John One", 1, "john_one"},
		{"John One", 2, "john_one_1"},
		{"John One", 3, "john_one_2"},
		{"John One Two", 1, "john_two"},
		{"John One Two", 2, "john_one"},
		{"John One Two", 3, "john_two_1"},
		{"John One Two", 4, "john_two_2"},
	}
	k := NewKname()
	line := 1
	for _, test := range tests {
		actual := k.GetShorten(test.name, test.attempt)
		if actual != test.expected {
			t.Errorf("Expected %v, got %v in line %v", test.expected, actual, line)
		}
		line++
	}
}
