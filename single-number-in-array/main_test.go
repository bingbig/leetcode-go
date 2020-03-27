package main

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	input := []int{3, 4, 3, 3}
	want := 4
	got := singleNumber(input)

	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
	input = []int{9, 1, 7, 9, 7, 9, 7}
	want = 1

	got = singleNumber(input)

	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}
