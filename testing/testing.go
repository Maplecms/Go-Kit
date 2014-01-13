package testing

import (
	"testing"
)

func If(t *testing.T, message string, expression bool) {
	if !expression {
		t.Error(message)
	}
}

func Unless(t *testing.T, message string, expression bool) {
	if expression {
		t.Error(message)
	}
}

func NoError(t *testing.T, e error) {
	if e != nil {
		t.Error(e.Error())
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
