package main

import (
	"testing"
)

func TestMinNumber(t *testing.T) {
	for i := 1; i < 3; i++ {
		minfunc := getFunc(i)
		got := minfunc([]int{10, 2})
		if "102" != got {
			t.Errorf("want 102 but got %s", got)
		}

		got = minfunc([]int{3, 30, 34, 5, 9})
		if "3033459" != got {
			t.Errorf("want 3033459 but got %s", got)
		}
	}
}

func getFunc(t int) func([]int) string {
	switch t {
	case 1:
		return minNumber
	case 2:
		return minNumber2
	}
	return nil
}
