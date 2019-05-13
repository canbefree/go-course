package error

import (
	"fmt"
	"path"
	"runtime"
)

type Error struct {
	code  int
	msg   string
	where string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code = %d ; msg = %s", e.code, e.msg)
}

func New(code int, msg string) *Error {
	where := caller(1)
	return &Error{code: code, msg: msg, where: where}
}

func Warp(err error, msg string) *Error {
	var where string
	var code int
	switch t := err.(type) {
	case *Error:
		// 继承where和code
		where = t.where
		code = t.code
		// 拼接上之前的错误
		msg = msg + ":: " + t.msg
	default:
		where = caller(1)
	}

	return &Error{code: code, msg: msg, where: where}
}

func caller(int) string {
	_, filename, line, ok := runtime.Caller(1)
	var info string
	if ok {
		info = fmt.Sprintf("%v:[%v]", path.Join(path.Dir(filename), ""), line) // the the main function file directory
	} else {
		info = "./"
	}
	return info
}
