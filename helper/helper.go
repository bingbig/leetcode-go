package helper

import "testing"

func AssertionSame(t *testing.T, actual, expect interface{}) {
	switch actual.(type) {
	case int:
		if actual.(int) != expect.(int) {
			t.Errorf("want %d, but got %d", expect.(int), actual.(int))
		}
	}
}
