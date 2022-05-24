package solution

import "testing"

func TestValidateStackSequences(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	popped2 := []int{4, 3, 5, 1, 2}

	if false == validateStackSequences(pushed, popped) {
		t.Errorf("want true but got false")
	}

	if true == validateStackSequences(pushed, popped2) {
		t.Errorf("want false but got true")
	}
}
