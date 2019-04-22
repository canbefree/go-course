package event

//定义消息类型

//MsgEvent 消息发送
type MsgEvent struct {
	uid int8
	msg string
}

//JoinEvent 通知时间
type JoinEvent struct {
}

type BoardCast struct {
}

type LeaveEvent struct {
}
