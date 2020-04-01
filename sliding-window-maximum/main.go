package main

import (
	"fmt"
)

type queue struct {
	Len  int
	Data []int
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}

	result := []int{}
	q := NewQueue()
	for i := 0; i < k; i++ {
		for !q.IsEmpty() && nums[q.Back()] <= nums[i] {
			q.PopBack()
		}
		q.PushBack(i)
	}

	for i := k; i < len(nums); i++ {
		result = append(result, nums[q.Front()])
		for !q.IsEmpty() && nums[q.Back()] <= nums[i] {
			q.PopBack()
		}
		if !q.IsEmpty() && q.Front() <= i-k {
			q.PopFront()
		}

		q.PushBack(i)
	}

	result = append(result, nums[q.Front()])
	return result
}

func NewQueue() *queue {
	return &queue{
		Len:  0,
		Data: make([]int, 0),
	}
}

func (q *queue) Back() int {
	return q.Data[q.Len-1]
}

func (q *queue) Front() int {
	return q.Data[0]
}

func (q *queue) IsEmpty() bool {
	return q.Len == 0
}

func (q *queue) PopBack() {
	q.Data = q.Data[:q.Len-1]
	q.Len--
}

func (q *queue) PopFront() {
	q.Data = q.Data[1:]
	q.Len--
}

func (q *queue) Clear() {
	q.Len = 0
	q.Data = make([]int, 0)
}

func (q *queue) PushBack(n int) {
	q.Data = append(q.Data, n)
	q.Len++
}

func main() {
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}
