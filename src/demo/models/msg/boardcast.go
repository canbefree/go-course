package msg

import (
	"errors"

	"github.com/gorilla/websocket"
)

//BoardCast 广播消息
type BoardCast struct {
	Msg
}

func (msg BoardCast) Handle(conn *websocket.Conn) error {
	conn.WriteMessage(websocket.TextMessage, []byte("msg from boardcast"))
	return errors.New("error")
}
