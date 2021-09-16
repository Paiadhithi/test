package main

import "fmt"

func main() {
	out := make(chan int)
	in := make(chan int)

	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)

	//All above func are waiting for values to be added in Input channe;l

	go func() {
		in <- 1
		in <- 2
		in <- 3
		in <- 4

	}()

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	//No need to sleep for using go routine

}
func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("Within goroutine")
	for {
		number := <-in
		r := number * 2
		out <- r
	}
}

