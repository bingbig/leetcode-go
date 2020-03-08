package main

import "testing"

func TestMajorityElement(t *testing.T) {
	t.Run("2,2,1,1,1,2,2", testMajorityElementFunc([]int{2, 2, 1, 1, 1, 2, 2}, 2))
	t.Run("1", testMajorityElementFunc([]int{1}, 1))
	t.Run("3,1,3", testMajorityElementFunc([]int{3, 1, 3}, 3))
}

func testMajorityElementFunc(arr []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		got := majorityElement(arr)
		if expected != got {
			t.Errorf("expect %d, but got %d", expected, got)
		}
	}
}
