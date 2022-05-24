package solution

import "container/heap"

type MinIntHeap struct {
	d []int
}

func (h MinIntHeap) Root() int {
	if h.Len() == 0 {
		return 0
	}
	return h.d[0]
}

func (h MinIntHeap) Len() int            { return len(h.d) }
func (h MinIntHeap) Less(i, j int) bool  { return h.d[i] < h.d[j] }
func (h MinIntHeap) Swap(i, j int)       { h.d[i], h.d[j] = h.d[j], h.d[i] }
func (h *MinIntHeap) Push(x interface{}) { (*h).d = append((*h).d, x.(int)) }
func (h *MinIntHeap) Pop() interface{} {
	old := (*h).d
	n := len(old)
	x := old[n-1]
	(*h).d = old[0 : n-1]
	return x
}

type MaxIntHeap struct {
	MinIntHeap
}

func (H MaxIntHeap) Less(i, j int) bool { return H.d[i] > H.d[j] }

type MedianFinder struct {
	maxHeap *MaxIntHeap
	minHeap *MinIntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	maxHeap := &MaxIntHeap{MinIntHeap{d: []int{}}}
	minHeap := &MinIntHeap{d: []int{}}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFinder{
		maxHeap: maxHeap,
		minHeap: minHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxHeap, num)
	heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
	if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(this.maxHeap, heap.Remove(this.minHeap, 0))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() == this.maxHeap.Len() {
		return float64(this.minHeap.Root()+this.maxHeap.Root()) * 0.5
	}

	return float64(this.maxHeap.Root())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
