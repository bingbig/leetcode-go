package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	want := 5
	got := maxProfit(prices)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}

	prices = []int{7, 6, 4, 3, 1}
	want = 0
	got = maxProfit(prices)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}
