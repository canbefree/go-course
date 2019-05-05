package protocol

//Protocol 协议
type Protocol interface {
	GetCMD() int
	Decode([]byte)
	Encode() ([]byte, error)
	GetContent() string //获取消息内容
	GetFriendID() int
}

//ClientProtocol 客户端协议
type ClientProtocol interface {
	Protocol
	GetFromID() int
}

//ServerProtocol 服务器协议
type ServerProtocol interface {
	Protocol
	//Handle 处理output消息格式推送给客户端
	Handle()
}

type Base struct {
	CMD     int    //指令
	FromID  int    //从哪里来
	ToID    int    //发给谁？
	Version string //版本号
	Content string //内容
}

func (b *Base) Decode([]byte) {}

func (b *Base) Encode() ([]byte, error) {
	return nil, nil
}

func (b *Base) GetCMD() int {
	return b.CMD
}

func (b *Base) GetContent() string {
	return b.Content
}

func (b *Base) GetFriendID() int {
	return b.ToID
}
