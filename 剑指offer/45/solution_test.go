package solution

import "testing"

func TestMinNumber(t *testing.T) {
	input := []int{10, 2}
	want := "102"
	got := minNumber(input)
	if want != got {
		t.Errorf("want %s but got %s", want, got)
	}

	input = []int{3, 30, 34, 5, 9}
	want = "3033459"
	got = minNumber(input)
	if want != got {
		t.Errorf("want %s but got %s", want, got)
	}
}
