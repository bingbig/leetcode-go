package main

import "testing"

func TestFindFirstNum(t *testing.T) {
	want := -1
	got := findFirstNum([]int{1, 2, 3, 4, 5}, 0, 0, 4)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}

	want = 3
	got = findFirstNum([]int{1, 2, 3, 4, 5}, 4, 0, 4)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}

	arr := []int{1, 2, 3, 4, 4, 4, 4, 4, 5}
	want = 3
	got = findFirstNum(arr, 4, 0, len(arr)-1)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}

	arr = []int{}
	want = -1
	got = findFirstNum(arr, 4, 0, len(arr)-1)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}

	arr = []int{2, 2}
	want = -1
	got = findFirstNum(arr, 3, 0, 1)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}

func TestFindLastNum(t *testing.T) {
	want := -1
	got := findLastNum([]int{1, 2, 3, 4, 5}, 0, 0, 4)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}

	want = 3
	got = findLastNum([]int{1, 2, 3, 4, 5}, 4, 0, 4)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}

	arr := []int{1, 2, 3, 4, 4, 4, 4, 4, 5}
	want = 7
	got = findLastNum(arr, 4, 0, len(arr)-1)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}

	arr = []int{}
	want = -1
	got = findLastNum(arr, 4, 0, len(arr)-1)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}
