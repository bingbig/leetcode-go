package main

import (
	"testing"
)

func TestFindNthDigit(t *testing.T) {
	input := []int{0, 1, 3, 10, 11, 10000}
	want := []int{0, 1, 3, 1, 0, 7}
	for i := 0; i < len(input); i++ {
		got := findNthDigit(input[i])
		if got != want[i] {
			t.Errorf("want: %d, but got: %d\n", want[i], got)
		}
	}

}
