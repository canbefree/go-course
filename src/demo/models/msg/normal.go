package msg

//Normal 广播消息
type Normal struct {
	Msg
}

func (msg Normal) SetResponse() error {
	msg.Response = "ss"
	return nil
}
