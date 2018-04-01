package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// If we don't close the channel, the channel try to wait and it lead deadlock...
	//
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan receive]:
	// main.main()
	// /tmp/sandbox585571777/main.go:19 +0x120
	//
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
