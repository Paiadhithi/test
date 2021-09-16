package main

import (
	"fmt"
)

//Before using channels, don't know when the go routine will finish execution
func main() {
	n := 2
	//Creating channel to send int data
	channel1 := make(chan int)
	//go multiplyByTwo(n) Using the channel
	go multiplyByTwo(n, channel1)

	//time.Sleep(time.Second) //Don't know when the GO routine is going to end
	fmt.Println(<-channel1) //When poutput is recievd on this chanenel, print it to the console
}

func multiplyByTwo(n int, out chan<- int) {
	result := n * 2
	out <- result

}

