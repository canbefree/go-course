package helper

import (
	"demo/helper"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {
	Convey("int", t, func() {
		a := 10
		original := reflect.ValueOf(a)
		typ := reflect.TypeOf(a)
		Convey("type", func() {
			So(original.Kind(), ShouldEqual, reflect.Int)
			So(typ.Kind(), ShouldEqual, reflect.Int)
			So(typ.Name(), ShouldEqual, "int")
		})
	})
}

func TestStruct(t *testing.T) {
	Convey("struct", t, func() {
		type Student struct {
			Name string
			Age  int
		}

		// a := &Student{Name: "小明", Age: 18}
		// original := reflect.ValueOf(a)

		// typ := reflect.TypeOf(a)
		Convey("type", func() {
			b := &Student{Name: "小光", Age: 19}
			originalb := reflect.ValueOf(b)
			orignStruct := originalb.Elem() //返回结构体
			src := reflect.New(originalb.Type()).Elem()
			src.Set(reflect.New(orignStruct.Type()))
			for i := 0; i < orignStruct.NumField(); i++ {
				field := orignStruct.Field(i)
				src.Elem().Field(i).Set(field)
			}
			So(src.Interface().(*Student), ShouldResemble, originalb.Interface().(*Student))
		})
	})
}

func TestDefine(t *testing.T) {
	Convey("test", t, func() {
		type Student struct {
			Name string
			Age  int
		}
		object := &Student{Name: "小明", Age: 18}

		Convey("define", func() {
			original := reflect.ValueOf(object)
			// Type() *Student
			So(original.Type().Kind(), ShouldEqual, original.Kind())
		})

		Convey("deep-copy", func() {
			subject := helper.DeepCopy(object)
			subject.(*Student).Age = 19
			So(subject, ShouldNotEqual, object)
		})

	})
}
