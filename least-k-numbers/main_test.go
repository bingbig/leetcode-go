package main

import "testing"

func TestHeapSort(t *testing.T) {
	cases := map[int][]int{
		1: []int{0, 0, 1, 0},
		2: []int{2},
		5: []int{5, 4, 3, 2, 1, 0},
	}

	for expect, test := range cases {
		heapSort(test)
		got := test[len(test)-1]
		if got != expect {
			t.Errorf("want %d but got %d", expect, got)
		}
	}
}

func TestGetLeastNumbers(t *testing.T) {
	data := []int{3, 2, 1}
	out := getLeastNumbers(data, 2)
	heapSort(out)
	if !(out[0] == 1 && out[1] == 2) {
		t.Errorf("failed")
	}
}

func TestQuickSort(t *testing.T) {
	cases := map[int][]int{
		1: []int{0, 0, 1, 0},
		2: []int{2},
		5: []int{5, 4, 3, 2, 1, 0},
	}

	for expect, test := range cases {
		quickSort(test, 0, len(test)-1)
		got := test[len(test)-1]
		if got != expect {
			t.Errorf("want %d but got %d", expect, got)
		}
	}
}
