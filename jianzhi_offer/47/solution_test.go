package solution

import "testing"

func TestMaxValue(t *testing.T) {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	want := 12
	got := maxValue(grid)
	if want != got {
		t.Errorf("want %d,  but got %d", want, got)
	}
}
