package main

import "fmt"

type Room struct {
	Users map[string]string
	Read  chan string //读消息通道
	Write chan string //写消息通道
}

func main() {
	for i := range []int{1, 2, 3, 4} {
		fmt.Println(i)
	}
}

func BoardCast(){
	for{
		select{
			case Write
		}
	}
}