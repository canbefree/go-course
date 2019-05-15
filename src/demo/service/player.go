package service

import (
	"demo/models/card"
)

type Player struct {
	card.Player
}

type IPlayer interface {
	Join(IServer) bool //加入服务器
}

func (player *Player) Join(server IServer) bool {
	server.AddPlayer(player)
	return true
}
