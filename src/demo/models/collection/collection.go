package collection

import (
	"demo/models/cmd"
	"demo/models/protocol"
	"demo/models/user"
	"log"
)

type Collection struct {
	//Users 所有的用户列表
	Users map[int]user.User
}

func NeCollection() *Collection {
	users := make(map[int]user.User)
	return &Collection{
		Users: users,
	}
}

func (s *Collection) Join(u user.User) {
	log.Printf("%v 加入服务器", u.ID)
	s.Users[u.ID] = u
}

func (s *Collection) Leave(u user.User) {
	log.Printf("%v 离开服务器", u.ID)
	delete(s.Users, u.ID)
}

func (s *Collection) BoardCast(msg string) {
	log.Printf("服务器广播消息 %v  %v", msg, s.Users)
	p := protocol.NewServer()
	p.CMD = cmd.UserBoardCast
	p.Content = msg
	for _, user := range s.Users {
		user.Output <- p
	}
}

func (s *Collection) sendMsg(uid int, msg string) {
	log.Printf("服务器给 %v 发送消息 %v", u.ID, msg)
	s.Users[u.ID].Output <- msg
}

//处理客户端线程的消息 转发线程
func (s *Collection) Handle(input chan protocol.ClientProtocol) {
	for {
		select {
		case p := <-input:
			switch p.GetCMD {
			case cmd.UserMessage:
				s.sendMsg(p.GetFid(), p.GetContent())
				break
			case cmd.UserBoardCast:
				s.BoardCast(p.GetContent())
				break
			}
		}
	}
}
