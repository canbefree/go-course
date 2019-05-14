package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Student struct {
	Name string
	Age  int
}

func TestDefine(t *testing.T) {
	Convey("test", t, func() {
		object := &Student{Name: "小明", Age: 18}
		Convey("deep-copy", func() {
			var subject Student
			DeepCopy(subject, object)
			subject.Age = 19
			So(subject, ShouldNotEqual, object)
		})
		Convey("define", func() {
			var s1 Student
			s2 := new(Student)
			s3 := &Student{}
			So(s2, ShouldResemble, s3)
			So(s1, ShouldNotResemble, s2)
		})
	})
}

func DeepCopy(dst interface{}, src interface{}) {
	// srcv := reflect.ValueOf(src)
	// srct := reflect.TypeOf(src)

	// dstv := reflect.ValueOf(dst)
	// dstt := reflect.TypeOf(dst)

}
