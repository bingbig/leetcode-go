package main

import (
	"testing"
)

func TestIsStraight(t *testing.T) {
	got := isStraight([]int{1, 2, 3, 4, 5})
	if got == false {
		t.Error("1 should be true but got false")
	}

	got = isStraight([]int{0, 0, 1, 2, 5})
	if got == false {
		t.Error("2 should be true but got false")
	}

	got = isStraight([]int{0, 0, 1, 2, 13})
	if got == true {
		t.Error("3 should be false but got true")
	}
}
