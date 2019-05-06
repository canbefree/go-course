package test

import "testing"

type User struct {
	Name string
}

func (u *User) SetName(s string) {
	u.Name = s
}
func TestStruct(t *testing.T) {

	//go会自动分析

	// *p 代表指向p的值 &p代表指针地址
	u := &User{"你好"}
	u.SetName("??")
	a := &u
	b := *a
	b.Name = "**"
	// *a.Name = "as"
	t.Logf("%v:%v:%v", u, &u, *a)

}
