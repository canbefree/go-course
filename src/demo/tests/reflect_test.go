package test

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValueOf(t *testing.T) {
	Convey("valueof", t, func() {
		Convey("int", func() {
			// reflect.Value类型比较
			a := 10
			s := 10
			So(reflect.ValueOf(a), ShouldNotEqual, reflect.ValueOf(s))
		})

		Convey("type value relation", func() {
			a := 10
			So(reflect.TypeOf(a), ShouldEqual, reflect.ValueOf(a).Type())
		})

		Convey("convert", func() {
			var num int = 10000
			value := reflect.ValueOf(num)
			convertValue := value.Interface().(int)
			So(convertValue, ShouldEqual, 10000)
		})
	})
}
