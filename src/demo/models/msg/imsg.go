package msg

import (
	"github.com/gorilla/websocket"
)

type iMsg interface {
	//获取CMD
	GetCMD() int
	//获取消息内容
	GetBody() string
	//处理消息
	Handle(*websocket.Conn) error
	//设置回应
	SetResponse() error
}
