package main

import (
	"testing"
)

type TestData struct {
	Input string
	Want  string
}

var tests = []TestData{
	TestData{"a", "a"},
	TestData{"", ""},
	TestData{"a good   example", "example good a"},
	TestData{"  hello world!  ", "world! hello"},
	TestData{"the sky is blue", "blue is sky the"},
}

func TestReverseWord(t *testing.T) {
	for idx, td := range tests {
		got := reverseWords(td.Input)
		if got != td.Want {
			t.Errorf("test_%d: want '%s', but got '%s'", idx, td.Want, got)
		}
	}
}
