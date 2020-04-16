package solution

import "testing"

func TestFindNthDigit(t *testing.T) {
	n := 3
	w := 3
	g := findNthDigit(n)
	if w != g {
		t.Errorf("%d: want %d, but got %d", n, w, g)
	}

	n = 11
	w = 0
	g = findNthDigit(n)
	if w != g {
		t.Errorf("%d: want %d, but got %d", n, w, g)
	}

	n = 10
	w = 1
	g = findNthDigit(n)
	if w != g {
		t.Errorf("%d: want %d, but got %d", n, w, g)
	}
}
