package user

import (
	"github.com/gorilla/websocket"
)

// User 定义
type User struct {
	ID   int8           //用户ID
	Conn websocket.Conn //客户端连接
}

// NewUser 返回一个新的用户
func NewUser(id int8) *User {
	return &User{
		ID: id,
	}
}

func (u *User) SendMsg(msg string) {

}
