package main

import (
	"testing"
)

func TestSumNums(t *testing.T) {
	if 6 != sumNums(3) {
		t.Errorf("error")
	}

	if 45 != sumNums(9) {
		t.Errorf("error")
	}
}
