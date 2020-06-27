package solution

import (
	"container/list"
)

func replaceElements(arr []int) []int {
	N := len(arr)
	if N == 0 {
		return nil
	}
	if N == 1 {
		return []int{-1}
	}

	stack := list.New()
	stack.PushBack(-1)

	for i := N - 1; i >= 0; i-- {
		if stack.Back().Value.(int) <= arr[i] {
			stack.PushBack(arr[i])
		}
	}
	for i := 0; i < N; i++ {
		if stack.Back().Value.(int) == arr[i] {
			stack.Remove(stack.Back())
		}

		arr[i] = stack.Back().Value.(int)
	}

	arr[N-1] = -1
	return arr
}
