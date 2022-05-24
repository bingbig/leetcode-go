package solution

import (
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	want := 2
	got := findRepeatNumber(nums)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}
