package pipe

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPipe(t *testing.T) {
	add := GetFristStack(add)

	Convey("t", t, func() {
	})
}

type NextClosure func(int) NextClosure
type NextFunc func(int, NextClosure) NextClosure

func add(num int, next NextClosure) NextClosure {
	return next(num)
}

func minues(num int, next NextClosure) NextClosure {
	return next(num)
}

func GetFristStack(des NextClosure) NextClosure {
	return func(num int) NextClosure {
		return des(num)
	}
}

func Call(_pipe NextFunc, stack NextClosure) NextClosure {
	return func(num int) NextClosure {
		return _pipe(num, stack)
	}
}
