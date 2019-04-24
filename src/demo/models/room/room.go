package room

import "os/user"

// Room 房间
type Room struct {
	Users []user.User
}
