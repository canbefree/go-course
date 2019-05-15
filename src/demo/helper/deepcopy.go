package helper

import "reflect"

//DeepCopy 深复制
//使用 DeepCopy(dst).(*int) 进行强制转换
func DeepCopy(src interface{}) interface{} {
	if src == nil {
		return nil
	}
	original := reflect.ValueOf(src)
	copy := reflect.New(original.Type()).Elem()

	copyRecursive(original, copy)

	return copy.Interface()
}

func copyRecursive(src, dst reflect.Value) {
	switch src.Kind() {
	case reflect.Ptr:
		original := src.Elem()
		newDst := reflect.New(original.Type())
		dst.Set(newDst) //将 dst的指针指向 newDst
		copyRecursive(original, newDst.Elem())
		break
	default:
		dst.Set(src)
	}
}
