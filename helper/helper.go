package helper

import (
	"reflect"
	"testing"
)

func AssertionSame(t *testing.T, got, expect interface{}) {
	switch got.(type) {
	case int:
		if got.(int) != expect.(int) {
			t.Errorf("expect %d, but got %d", expect.(int), got.(int))
		}
	case []int:
		if !reflect.DeepEqual(got.([]int), expect.([]int)) {
			t.Errorf("expect %v, but got %v", expect.([]int), got.([]int))
		}
	}
}
