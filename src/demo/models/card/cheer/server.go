package cheer

import "demo/models/user"

//Server 干瞪眼服务器算法
type Server struct {
	Players map[int]user.User
}

type IServer interface {
	Start() bool           //由房主触发
	AddPlayer(Player) bool //添加玩家
}
