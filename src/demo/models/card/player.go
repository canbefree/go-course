package card

import "demo/models/user"

type Player struct {
	user   user.User
	status int //玩家状态
	pos    POS //玩家位置
}
