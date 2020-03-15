package main

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	want := 6
	got := maxSubArray(arr)
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}

	arr = []int{1}
	got = maxSubArray(arr)
	if 1 != got {
		t.Errorf("want: 1, but got: %d", got)
	}

	arr = []int{}
	got = maxSubArray(arr)
	if 0 != got {
		t.Errorf("want: 0, but got: %d", got)
	}

	arr = []int{1, 2}
	got = maxSubArray(arr)
	if 3 != got {
		t.Errorf("want: 3, but got: %d", got)
	}
}
