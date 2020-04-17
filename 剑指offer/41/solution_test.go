package solution

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	mh := &MinIntHeap{d: []int{2, 1, 5, 4, 9}}
	heap.Init(mh)
	heap.Push(mh, 10)
	out := []int{}
	for mh.Len() > 0 {
		out = append(out, heap.Pop(mh).(int))
	}
	if !reflect.DeepEqual(out, []int{1, 2, 4, 5, 9, 10}) {
		t.Errorf("min int heap failed: %v", out)
	}

	mH := &MaxIntHeap{MinIntHeap{d: []int{2, 1, 5, 4, 9}}}
	heap.Init(mH)
	heap.Push(mH, 10)
	out = []int{}
	for mH.Len() > 0 {
		out = append(out, heap.Pop(mH).(int))
	}
	if !reflect.DeepEqual(out, []int{10, 9, 5, 4, 2, 1}) {
		t.Errorf("max int heap failed: %v", out)
	}
}

func TestFindMedian(t *testing.T) {
	mf := Constructor()
	mf.AddNum(2)
	mf.AddNum(3)
	if mf.FindMedian() != 2.5 {
		t.Error("median is not 2.5")
	}
	mf.AddNum(4)
	if mf.FindMedian() != 3 {
		t.Error("median is not 3")
	}
}
