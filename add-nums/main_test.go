package main

import "testing"

func TestAdd(t *testing.T) {
	if 0 != add(0, 0) {
		t.Error("error 0, 0")
	}
	if -1 != add(0, -1) {
		t.Error("error 0,-1")
	}

	if -4 != add(-2, -2) {
		t.Error("error -2, -2")
	}

	if 4 != add(1, 3) {
		t.Error("error 1,3")
	}
}
