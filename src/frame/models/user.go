package models

// type User struct {
// 	name string
// 	age  int8
// }

// type UserInterface interface {
// 	Name()
// }

type User struct {
	ID       string
	Username string
	Password string
}

func (u User) Name() string {
	return u.Username
}

func init() {
	u := User{"小米", "", ""}
	u.Name()
}
