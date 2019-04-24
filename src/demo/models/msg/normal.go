package msg

import (
	"errors"

	"github.com/gorilla/websocket"
)

//Normal 广播消息
type Normal struct {
	Msg
}

func (msg Normal) Handle(conn *websocket.Conn) error {
	conn.WriteMessage(websocket.TextMessage, []byte("msg from person"))
	return errors.New("msg")
}
