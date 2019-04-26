package user

import (
	"demo/models/protocol"
	"sync"

	"github.com/gorilla/websocket"
)

// User 定义
type User struct {
	ID     int                          //用户ID
	Conn   websocket.Conn               //客户端连接
	Input  chan protocol.ClientProtocol //输入数据
	Output chan protocol.ServerProtocol //输出数据
}

//输出
type Output struct {
}

// NewUser 返回一个新的用户
func NewUser(id int) *User {
	return &User{
		ID: id,
	}
}

func (u *User) Handle(wg *sync.WaitGroup, conn *websocket.Conn) {
	defer wg.Done()
	defer conn.Close()

	// go listen()?  通过chan来获取大厅的数据？
	go func() {
		for {
			select {
			case p := <-u.Output:

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

//HandleClientMsg 处理客户端的消息
func (u *User) HandleClientMsg(p protocol.Protocol) {
	u.Input <- p
}
