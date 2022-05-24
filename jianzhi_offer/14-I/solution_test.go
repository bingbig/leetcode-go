package solution

import "testing"

func TestCuttingRope(t *testing.T) {
	want := 86093442
	got := cuttingRope(50)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}

func TestCuttingRope2(t *testing.T) {
	want := 86093442
	got := cuttingRope2(50)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}

func BenchmarkCuttingRope(t *testing.B) {
	for i := 0; i < 100; i++ {
		cuttingRope(i)
	}
}

func BenchmarkCuttingRope2(t *testing.B) {
	for i := 0; i < 100; i++ {
		cuttingRope2(i)
	}
}
