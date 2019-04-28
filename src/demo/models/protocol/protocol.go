package protocol

//Protocol 协议
type Protocol interface {
	GetCMD() int
	Decode([]byte)
	Encode() ([]byte, error)
	GetContent() string //获取消息内容
	GetFID() int
}

//ClientProtocol 客户端协议
type ClientProtocol interface {
	Protocol
}

//ServerProtocol 服务器协议
type ServerProtocol interface {
	Protocol
	//Handle 处理output消息格式推送给客户端
	Handle()
}

type Base struct {
	CMD     int    //指令
	FID     int    //指定人ID
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

func (b *Base) GetFID() int {
	return b.FID
}
