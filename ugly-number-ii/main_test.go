package main

import (
	"testing"
)

func TestMinInt(t *testing.T) {
	got := minInt(1, 1, 1)
	if got != 1 {
		t.Errorf("want 1 but got %d", got)
	}

	got = minInt(2, 3, 1)
	if got != 1 {
		t.Errorf("want 1 but got %d", got)
	}

	got = minInt(100, 22, 111)
	if got != 22 {
		t.Errorf("want 22 but got %d", got)
	}
}

func TestNthUglyNumber(t *testing.T) {
	got := nthUglyNumber(0)
	if got != 0 {
		t.Errorf("want 0 but got %d", got)
	}

	got = nthUglyNumber(3)
	if got != 3 {
		t.Errorf("want 3 but got %d", got)
	}

	got = nthUglyNumber(10)
	if got != 12 {
		t.Errorf("want 12 but got %d", got)
	}

	got = nthUglyNumber(11)
	if got != 15 {
		t.Errorf("want 15 but got %d", got)
	}

}
