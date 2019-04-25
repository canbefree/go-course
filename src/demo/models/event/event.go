package event

//定义消息类型

const (
	TYPE_LEAVE = iota
	TYPE_JOIN
	TYPE_BOARDCAST
	TYPE_NOMAL
)

//MsgEvent 消息发送
type Event struct {
	TYPE int
	BODY string
}

type iEvent interface {
	GetType()
	GetBODY()
	Handle()
}
