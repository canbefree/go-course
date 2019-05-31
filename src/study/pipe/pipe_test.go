package pipe

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNormal(t *testing.T) {
	Convey("t", t, func() {
		result := add(minues(50)())()
		So(result, ShouldEqual, 10)
	})

	Convey("a", t, func() {
		num := 50
		for _, f := range [...](func(int) func() int){add, minues} {
			num = f(num)()
		}
		So(num, ShouldEqual, 10)
	})

}

func TestPipe(t *testing.T) {
	Convey("t", t, func() {

	})
}

func add(num int) func() int {
	params := 10
	return func() int {
		return num + params
	}
}

func minues(num int) func() int {
	params := 50
	return func() int {
		return num - params
	}
}
