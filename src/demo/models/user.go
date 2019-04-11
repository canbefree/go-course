package models


type User struct {
	Name string
	Age  int8
}

func (u *User) GetName() string {
	return u.Name
}
