package main

import "fmt"

func main() {
	ch := make(chan int, 2) // if the value is 1, then ch can only buffer one value.
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
