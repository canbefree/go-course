package test

import "testing"

func TestChan(t *testing.T) {
	out := make(chan int)
	go func() {
		for _, n := range [...]int{1, 3, 4, 5, 6, 7} {
			out <- n
		}
		close(out)
	}()

}
