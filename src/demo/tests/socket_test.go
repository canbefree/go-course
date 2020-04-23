package test

import (
	"net"
	"testing"
)

func TestServerEpoll(t *testing.T) {
	t.Logf("start")
	startServer(t)
}

func startServer(t *testing.T) {

	ch := make(chan net.Conn)

	ln, err := net.Listen("tcp", ":8902")
	if err != nil {
		t.Errorf("error:%v", err)
	}

	//处理客户端发来的消息
	for i := 0; i < 5; i++ {
		go func() {
			for c := range ch { //当有连接进来是就会循环 否则是一个空的轮询
				defer c.Close()
				for {
					buf := make([]byte, 1024)
					rmsg, err := c.Read(buf)
					if err != nil {
						t.Logf("erro")
						return
					}
					t.Logf("msg:%v", rmsg)

					_, err = c.Write(buf)
					if err != nil {
						t.Logf("erro")
					}
					t.Logf("send")
					return
				}
			}
		}()
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			t.Errorf("error:%v", err)
		}
		//将连接加入连接池，为什么用管道是因为只有管道在协程中共享
		ch <- conn
	}
}
