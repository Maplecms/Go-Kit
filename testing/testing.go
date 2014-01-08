package testing

import (
	"testing"
	"errors"
	"fmt"
)

type Tplus struct {
	testing.T
}

func New(t *testing.T) *Tplus {
	return &Tplus{*t}
}

func (self *Tplus) NoError(e error) {
	if e != nil {
		self.Error(e.Error())
	}
}

func Equal(x, y, interface {}) error {
	if x != y {
		return errors.New(fmt.Sprintf("%v != %v", x, y))
	}
	return nil
}

func Inequal(x, y, interface{}) error {
	if x == y {
		return errors.New(fmt.Sprintf("%v == %v", x, y))
	}
	return nil
}

func SliceEqual(xs, ys, []interface{}) error {
	if len(xs) != len(ys) {
		return errors.New(fmt.Sprintf("len(%v) != len(%v)", xs, ys))
	}
	for index, x := range xs {
		y := ys[index]
		if err := Equal(x, y); err != nil {
			return err
		}
	}
}











