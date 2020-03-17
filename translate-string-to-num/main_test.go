package main

import (
	"testing"
)

func TestTranslateNum(t *testing.T) {
	got := translateNum(12258)
	want := 5
	if got != want {
		t.Errorf("want %d, but got %d", want, got)
	}
}
