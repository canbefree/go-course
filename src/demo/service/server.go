package service

import (
	"sync"
)

type Server struct {
	Players []*Player
	lock    sync.Mutex
}

//服务器线程
type IServer interface {
	AddPlayer(player *Player) bool
}

func (server *Server) AddPlayer(player *Player) bool {
	server.lock.Lock()
	server.Players = append(server.Players, player)
	server.lock.Unlock()
	return true
}
