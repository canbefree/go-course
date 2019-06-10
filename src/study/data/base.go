package data

import "reflect"

//基础的数据结构有 int char float double

func GetType(d interface{}) interface{} {
	return reflect.TypeOf(d)
}
