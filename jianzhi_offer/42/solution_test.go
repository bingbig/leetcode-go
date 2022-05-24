package solution

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	want := 6
	got := maxSubArray(nums)

	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}
