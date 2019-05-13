package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRange(t *testing.T) {
	Convey("test range", t, func() {
		Convey("int", func() {
			var r []int
			// ret = []int{1, 2, 3, 4}
			r = []int{1, 2, 3, 4}
			var ret []int
			for k, _ := range r {
				t.Log(k)
				ret = append(ret, k)
			}
			So([]int{0, 1, 2, 3}, ShouldResemble, ret)
		})
	})
}
