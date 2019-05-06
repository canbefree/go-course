package user

import (
	"demo/models/cmd"
	"demo/models/protocol"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// User 定义
type User struct {
	ID     int                          //用户ID
	Conn   *websocket.Conn              //客户端连接
	Input  chan protocol.ClientProtocol //这里指客户端发来的消息
	Output chan protocol.ServerProtocol //这里指线程需要发给客户端的消息

	UniqueID string //全局唯一的ID
}

//输出
type Output struct {
}

// NewUser 返回一个新的用户
func NewUser(id int) *User {
	return &User{
		ID:       id,
		UniqueID: strconv.FormatInt(time.Now().Unix(), 10),
	}
}

// CollectMsg 收集消息
func (u *User) CollectMsg(p protocol.ServerProtocol) {
	u.Output <- p
}

func (u *User) Handle(wg *sync.WaitGroup) {
	conn := u.Conn
	defer wg.Done()
	defer u.deferHandle()

	//通过chan来获取大厅的数据？
	go func() {
		for {
			select {
			case p := <-u.Output:
				log.Printf("向客户端写入消息:%v", p)
				u.HandleMessageFromServer(p)
				break
			}
		}
	}()

	for {
		//一直读取客户端的消息
		contentType, content, err := conn.ReadMessage()

		if err != nil {
			conn.WriteMessage(contentType, []byte("error!"))
			break
		} else {
			switch contentType {
			case websocket.TextMessage:
				// 	//格式化message
				p := &protocol.Client{}
				p.Decode([]byte(content))
				log.Printf("服务器接收到消息:%v", p)
				u.HandleClientMsg(p)
				break
			case websocket.CloseMessage:
				conn.WriteMessage(contentType, []byte("byebye"))
				conn.Close()
				break
			}
		}
	}

}

func (u *User) deferHandle() {
	log.Printf("客户端退出:%v", u.ID) //客户端退出需要通知服务器
	p := &protocol.Client{}
	p.CMD = cmd.Leave
	p.FromID = u.ID
	p.Content = u.UniqueID
	u.Input <- p
	u.Close()
}

//HandleClientMsg 处理客户端的消息
func (u *User) HandleClientMsg(p protocol.ClientProtocol) {
	u.Input <- p
}

func (u *User) HandleMessageFromServer(p protocol.ServerProtocol) {
	u.Conn.WriteJSON(p)
	if p.GetCMD() == cmd.Kicking {
		u.Close()
	}
}

func (u *User) Close() {
	log.Printf("关闭连接:%v", u)
	u.Conn.Close()
}
