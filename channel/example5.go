package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan int)
	out2 := make(chan int)

	go fast(2, out1)
	go slow(3, out2)

	select { //whichever channel recieves the data first, will call the func
	case res := <-out1:
		fmt.Println("Fast finished first, result:", res)
	case res := <-out2:
		fmt.Println("slow finished first, result", res)
	}

}
func fast(n int, out chan<- int) {
	result := n * 2
	time.Sleep(5 * time.Millisecond)
	out <- result
}
func slow(n int, out chan<- int) {
	result := n * 2
	time.Sleep(15 * time.Millisecond)
	out <- result
}

