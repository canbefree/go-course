package main

/*
keyword
Pubsub 调度中心
User 客户端
Room	房间号


*/
import (
	"log"
	"time"
)

//订阅
type Pubsub interface {
	GetCollect() chan string
	Register(Observer) //注册
	Publish(string)    //发布消息
}

//观察者接口
type Observer interface {
	GetId() int8
	GetPipe() chan string //读消息通道
	Subscripe(Pubsub)     //订阅消息
	Listen()              //监听消息
	setWrite(chan string) //设置订阅的写入
	Publish(string)       //广播消息
}

type User struct {
	Id    int8
	Name  string
	Write chan string
	Pipe  chan string
}

func (u *User) Publish(msg string) {
	u.Write <- msg
}

func (u *User) Listen() {
	for {
		time.Sleep(100 * time.Microsecond)
		select {
		case msg := <-u.Pipe:
			log.Println("from room: %v", msg)
		default:
			break
		}
	}
}

func (u *User) GetId() int8 {
	return u.Id
}

func (u *User) GetPipe() chan string {
	return u.Pipe
}

func (u *User) Subscripe(s Pubsub) {
	u.Write = s.GetCollect()
}

func (u *User) setWrite(ch chan string) {
	u.Write = ch
}

func NewUser(id int8, name string) Observer {
	ch := make(chan string)
	pipe := make(chan string)
	return &User{
		Id:    id,
		Name:  name,
		Write: ch,
		Pipe:  pipe,
	}
}

type Room struct {
	RoomId  int8
	Users   map[int8]chan string
	Collect chan string
}

func (r *Room) Register(u Observer) {
	r.Users[u.GetId()] = u.GetPipe()
	u.setWrite(r.GetCollect())
}

func (r *Room) GetCollect() chan string {
	return r.Collect
}

//广播消息
func (r *Room) Publish(msg string) {
	for _, ch := range r.Users {
		log.Println("send to %v", r.Users)
		ch <- msg
		log.Println("end send to %v", r.Users)
	}
}

func NewRoom(rid int8) Pubsub {
	return &Room{
		RoomId: rid,
		Users:  make(map[int8]chan string),
	}
}

func main() {
	room := NewRoom(1) //建立一个room唯一的频道
	user1 := NewUser(1, "小明")
	user2 := NewUser(2, "小白")

	room.Register(user1)
	room.Register(user2)

	go user1.Listen() //user1监听消息
	go user2.Listen() //User2监听消息

	for {
		time.Sleep(time.Second)
		log.Println("?????")
		room.Publish("hello,somebody??")
	}

}
