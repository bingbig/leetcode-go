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
	case float32:
		if got.(float32) != expect.(float32) {
			t.Errorf("expect %f, but got %f", expect.(float32), got.(float32))
		}
	case []int:
		if !reflect.DeepEqual(got.([]int), expect.([]int)) {
			t.Errorf("expect %v, but got %v", expect.([]int), got.([]int))
		}
	case bool:
		if got.(bool) != expect.(bool) {
			t.Errorf("expect %v, but got %v", expect.(bool), got.(bool))
		}
	case string:
		if got.(string) != expect.(string) {
			t.Errorf("expect %v, but got %v", expect.(string), got.(string))
		}
	default:
		t.Errorf("Unexpected type\n")
	}
}
