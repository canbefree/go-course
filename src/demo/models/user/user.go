package user

// User 定义
type User struct {
	ID   int8   //用户ID
	Name string //用户姓名
}

// NewUser 返回一个新的用户
func NewUser(token string) *User {
	id := len(token)
	return &User{
		ID:   int8(id),
		Name: string(token),
	}
}
