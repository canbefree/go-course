package main

import "time"

func main() {

	test := make(chan int)
	go func() {
		for {
			time.Sleep(time.Second)
			test <- 123
		}
	}()

	for {
		select {
		case ch := <-test:
			println(ch)
		default:
			println("timeout")
		}
	}
}
