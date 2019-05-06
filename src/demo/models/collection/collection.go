package collection

import (
	"demo/models/cmd"
	"demo/models/protocol"
	"demo/models/user"
	"log"
	"time"
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
	//检查当前连接是否在线,是否需要踢出登录用户
	existUser, ok := s.Users[u.ID]
	//定义一个锁
	if ok == true {
		//todo 踢人 我踢我自己?
		s.Kicking(existUser)
		// existUser.Close()
		log.Printf("踢人成功! %v被踢 -- 当前用户列表:%v", u.ID, s.GetOnlineUsersID())
	}
	//覆盖掉之前的连接
	s.Users[u.ID] = u
	log.Printf("%v 加入服务器 %v", u.ID, s.GetOnlineUsersID())
}

func (s *Collection) Kicking(u user.User) {
	log.Printf("踢人")
	p := protocol.NewServer()
	p.CMD = cmd.Kicking
	p.Content = "你被踢了,不好意思哈" + time.Now().String()
	u.CollectMsg(p)
	// 由于是异步请求，不能保证用户列表原子性。所以由踢人改成覆写
	// delete(s.Users, u.ID)
}

func (s *Collection) Leave(uniqueID string) {
	for k, v := range s.Users {
		if v.UniqueID == uniqueID {
			delete(s.Users, k)
			log.Printf("%v 离开服务器 %v", k, s.GetOnlineUsersID())
			return
		}
	}
	// delete(s.Users, userId)
}

func (s *Collection) BoardCast(msg string) {
	log.Printf("服务器广播消息 %v  %v", msg, s.Users)
	p := protocol.NewServer()
	p.CMD = cmd.UserBoardCast
	p.Content = msg
	for _, user := range s.Users {
		user.CollectMsg(p)
	}
}

func (s *Collection) SendMsg(uid int, msg string) {
	log.Printf("服务器给 %v 发送消息 %v", uid, msg)
	p := protocol.NewServer()
	p.CMD = cmd.UserMessage
	p.Content = msg
	user, ok := s.Users[uid]
	if ok {
		user.CollectMsg(p)
	}
}

func (s *Collection) GetOnlineUsersID() []int {
	var usersId []int
	for k, _ := range s.Users {
		usersId = append(usersId, k)
	}
	return usersId
}

//处理客户端线程的消息 转发线程
func (s *Collection) Handle(input chan protocol.ClientProtocol) {
	for {
		select {
		case p := <-input:
			switch p.GetCMD() {
			case cmd.UserMessage:
				s.SendMsg(p.GetFriendID(), p.GetContent())
				break
			case cmd.UserBoardCast:
				s.BoardCast(p.GetContent())
				break
			case cmd.Leave:
				// uid := p.GetFromID()
				uniqueID := p.GetContent()
				s.Leave(uniqueID)
				break
			}
		}
	}
}
