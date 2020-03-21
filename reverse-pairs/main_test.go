package main

import "testing"

func TestReversePaires(t *testing.T) {
	arr, count := []int{1, 3, 2, 3, 1}, 4
	got := reversePairs(arr)

	if got != count {
		t.Errorf("want %d, but got %d", count, got)
	}

	arr, count = []int{7, 5, 6, 4}, 5
	got = reversePairs(arr)

	if got != count {
		t.Errorf("want %d, but got %d", count, got)
	}

	arr, count = []int{}, 0
	got = reversePairs(arr)

	if got != count {
		t.Errorf("want %d, but got %d", count, got)
	}

	arr, count = []int{1, 2}, 0
	got = reversePairs(arr)

	if got != count {
		t.Errorf("want %d, but got %d", count, got)
	}

	arr, count = []int{1, 2, 3, 4, 5, 6}, 0
	got = reversePairs(arr)

	if got != count {
		t.Errorf("want %d, but got %d", count, got)
	}
}
