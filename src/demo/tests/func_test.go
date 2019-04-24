package test

import "testing"

type User struct {
	Name string
}

func (u User) SetName(s string) {
	u.Name = s
}
func TestStruct(t *testing.T) {
	u := &User{"你好"}
	a.SetName("???")
	t.Logf("%v", a)
}
