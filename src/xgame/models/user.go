package models

import "github.com/gorilla/websocket"

type User struct {
	uid  int
	conn websocket.Conn
}

var UserList map[int64]chan Event
