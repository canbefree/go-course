package user

import (
	"demo/models/msg"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// User 定义
type User struct {
	ID     int            //用户ID
	Conn   websocket.Conn //客户端连接
	Input  chan string    //输入数据
	Output chan string    //输出数据
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
			case msgFromOtherUser := <-u.Output:
				conn.WriteMessage(websocket.TextMessage, []byte(msgFromOtherUser))
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
				message, err := msg.JSONDecode(string(content))
				if err != nil {
					log.Printf("err: %v", err)
					break
				}
				fmt.Println(message)
				u.Input <- message.GetBody()
				log.Printf("接受消息 %v", message.GetBody())
				// u.Input
				// message.Handle(conn)
				break
			case websocket.CloseMessage:
				conn.WriteMessage(contentType, []byte("byebye"))
				conn.Close()
				break
			}
		}
	}

}
