package helper

import (
	"demo/helper"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//内存存储是不连续的

// 赋值一个指针变量
// 变量地址 内容~结构地址
// 结构地址 内容~结构内容

//赋值一个整形
// 整形地址 内容~整形内容

func TestDeepCopy(t *testing.T) {
	Convey("deep copy", t, func() {
		Convey("int", func() {
			//整形赋值
			i := 10
			a := i
			b := helper.DeepCopy(i).(int)
			a = 1

			So(i, ShouldNotEqual, a)
			So(b, ShouldEqual, 10)
		})
		Convey("struct", func() {
			type Student struct {
				Name string
				Age  int
			}
			student := &Student{Name: "小明", Age: 18}
			student_copy := helper.DeepCopy(student).(*Student)
			So(student, ShouldResemble, student_copy)
			student_copy.Name = "小光"
			So(student_copy, ShouldResemble, &Student{Name: "小光", Age: 18})
		})
	})
}

func TestUse(t *testing.T) {
	Convey("define", t, func() {
		type Student struct {
			Name string
			Age  int
		}

		object := &Student{Name: "小明", Age: 18}

		Convey("Type", func() {
			a := 10
			original := reflect.ValueOf(a)
			typ := reflect.TypeOf(a)
			Convey("type", func() {
				// Kind()返回是一个Type类型 ： reflect.Int
				So(original.Kind(), ShouldEqual, reflect.Int)
				// 变量的 类型和值对象 Kind() 方法返回相同
				So(typ.Kind(), ShouldEqual, reflect.Int)
				// 类型可以转换为string
				So(typ.Name(), ShouldEqual, "int")
			})
		})

		Convey("use reflect modify object", func() {
			b := &Student{Name: "小光", Age: 19}
			originalb := reflect.ValueOf(b)
			orignStruct := originalb.Elem()             //Elem()相当于 *ptr
			src := reflect.New(originalb.Type()).Elem() // 创建一个 &Student 然后返回他的指针指向的结构地址
			src.Set(reflect.New(orignStruct.Type()))
			for i := 0; i < orignStruct.NumField(); i++ {
				field := orignStruct.Field(i)
				src.Elem().Field(i).Set(field)
			}
			So(src.Interface().(*Student), ShouldResemble, originalb.Interface().(*Student))
		})

		Convey("define", func() {
			original := reflect.ValueOf(object)
			// Type() *Student
			So(original.Type().Kind(), ShouldEqual, original.Kind())
		})

	})
}
