package solution

import "testing"

func TestPushMin(t *testing.T) {
	ms := Constructor()
	top := 0
	min := 0
	ms.Push(1)
	min = ms.Min()
	top = ms.Top()
	if top != 1 && min != 1 {
		t.Error("failed")
	}
	ms.Push(2)
	top = ms.Top()
	min = ms.Min()
	if top != 2 && min != 1 {
		t.Error("failed")
	}
	ms.Push(3)
	top = ms.Top()
	min = ms.Min()
	if top != 3 && min != 1 {
		t.Error("failed")
	}
	ms.Push(4)
	top = ms.Top()
	min = ms.Min()
	if top != 4 && min != 1 {
		t.Error("failed")
	}
	ms.Push(-1)
	top = ms.Top()
	min = ms.Min()
	if top != 4 && min != -1 {
		t.Error("failed")
	}
}
