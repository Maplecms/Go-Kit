package testing

import (
	"testing"
	"reflect"
)

func Expect(t *testing.T, message string, expression bool) {
	if !expression {
		t.Error(message)
	}
}

func NoError(t *testing.T, e error) {
	if e != nil {
		t.Error(e.Error())
	}
}

func DeepEqual(t *testing.T, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		t.Errorf("%v != %v", x, y)
	}
}

func Equal(t *testing.T, x, y interface {}) {
	if x != y {
		t.Errorf("%v != %v", x, y)
	}
}

func Inequal(t *testing.T, x, y interface{}) {
	if x == y {
		t.Errorf("%v == %v", x, y)
	}
}

func Contains(t *testing.T, xs []interface{}, y interface{}) {
	for _, x := range xs {
		if x == y {
			return
		}
	}
	t.Errorf("%v is not an element of %v", y, xs)
}

func ContentsEqual(t *testing.T, xs, ys, []interface{}) {
	for _, x := range xs {
		Contains(t, ys, x)
	}
}

func BytesEqual(t *testing.T, xs, ys []byte) {
	if len(xs) != len(ys) {
		t.Errorf("len(%v) != len(%v)", xs, ys)
	}
	for index, x := range xs {
		Equal(t, x, ys[index])
	}
}

func BytesInequal(t *testing.T, xs, ys []byte) {
	if len(xs) == len(ys) {
		for index, x := range xs {
			Equal(t, x, ys[index])
		}
	}
}
