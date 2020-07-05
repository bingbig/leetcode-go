package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{43, 65, 879, 834, 232, 0, 5446, 2, 0, 1, 2, 3, 53, 2, 1, 32, 435, 46, 56}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestQuickSort2(t *testing.T) {
	nums := []int{43, 65, 879, 834, 232, 0, 5446, 2, 0, 1, 2, 3, 53, 2, 1, 32, 435, 46, 56}

	fmt.Println(QuickSort2(nums))
}

func TestQuickSort3(t *testing.T) {
	nums := []int{43, 65, 879, 834, 232, 0, 5446, 2, 0, 1, 2, 3, 53, 2, 1, 32, 435, 46, 56}
	QuickSort3(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestBubbleSort(t *testing.T) {
	nums := []int{43, 65, 879, 834, 232, 0, 5446, 2, 0, 1, 2, 3, 53, 2, 1, 32, 435, 46, 56}
	BubbleSort(nums)
	fmt.Println(nums)
}

func TestInsertSort(t *testing.T) {
	nums := []int{43, 65, 879, 834, 232, 0, 5446, 2, 0, 1, 2, 3, 53, 2, 1, 32, 435, 46, 56}
	InsertSort(nums)
	fmt.Println(nums)
}

func TestShellSort(t *testing.T) {
	nums := []int{43, 65, 879, 834, 232, 0, 5446, 2, 0, 1, 2, 3, 53, 2, 1, 32, 435, 46, 56}
	shellSort(nums)
	fmt.Println(nums)
}
