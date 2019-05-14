package test

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAppend(t *testing.T) {
	Convey("array", t, func() {
		arr1 := []int{1, 2, 3, 4, 5}
		arr2 := make([]int, 3)
		Convey("merge1", func() {
			arr1 = append(arr1, arr2...)
			except := []int{1, 2, 3, 4, 5, 0, 0, 0}
			So(arr1, ShouldResemble, except)
		})

		Convey("copy", func() {
			//第一个参数 为目标数组
			copy(arr1, arr2)
			except := []int{0, 0, 0, 4, 5}
			So(arr1, ShouldResemble, except)
		})

	})
}

func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

func TestArrayAppend(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	slice := arr[1:3]

	slice = append(slice, 1)

	except := []int{2, 3, 1}

	for i := range slice {
		if slice[i] != except[i] {
			t.Errorf("slice not equal except: %v %v", slice, except)
		}
	}
}

func TestArray(t *testing.T) {
	convey.Convey("testArray", t, func() {
		a := []int{1, 2, 3}
		convey.Convey("index", func() {
			a[1] = 3
			convey.So(a, convey.ShouldEqual, [...]int{3})
		})
	})
}
