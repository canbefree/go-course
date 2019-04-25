package msg

//BoardCast 广播消息
type BoardCast struct {
	Msg
}

func (msg *BoardCast) SetResponse() error {
	msg.Response = ""
	return nil
}
