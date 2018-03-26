package main

import "fmt"

func fibonacci() func() int {
	x1, x2 := 0, 1
	return func() int {
		answer := x1
		x1 = x2
		x2 = answer + x1
		return answer
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
