package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, float64(2), FindMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.5, FindMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
