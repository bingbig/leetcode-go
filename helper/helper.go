package helper

import "testing"

func AssertionSame(t *testing.T, got, expect interface{}) {
	switch got.(type) {
	case int:
		if got.(int) != expect.(int) {
			t.Errorf("expect %d, but got %d", expect.(int), got.(int))
		}
	}
}
