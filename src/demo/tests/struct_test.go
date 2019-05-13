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
	Convey("l", t, func() {
		var s1 Student
		s1.Age = 123
		s2 := new(Student)
		s3 := &Student{}
		So(s2, ShouldResemble, s3)
		So(s1, ShouldResemble, s2)
	})
}
