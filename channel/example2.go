package main

import (
	"fmt"
)

func main() {
	var ch chan int
	fmt.Printf("%v %v %v", ch, len(ch), cap(ch))
	//ch <- 4  //Sending the data
	//<-ch    //Recieving the dat , should no tbe nil
	//Making unbuffered channel using make
	ch = make(chan int)
	//No capacity to queue
	fmt.Printf("Unbuffered :ch:%v, len: %v, cap: %v \n", ch, len(ch), cap(ch))
	//ch <- 2
	//Reciever should always be ready to send
	//<-ch

	//Creating a buffered channel ie with the capacity
	ch = make(chan int, 1)
	fmt.Printf("Buffered : ch:%v, len:%v, cap:%v\n", ch, len(ch), cap(ch))
	ch <- 2
	fmt.Printf("Buffered : ch:%v, len:%v, cap:%v\n", ch, len(ch), cap(ch))
	v := <-ch //removed from channel
	fmt.Printf("Recieved from channel %v\n", v)

	//Channel with more capacity
	ch = make(chan int, 2)
	ch <- 2
	ch <- 3
	//recieving from an empty channels
	fmt.Printf("Buffered : ch:%v, len:%v, cap:%v\n", <-ch, len(ch), cap(ch))
	fmt.Printf("Buffered : ch:%v, len:%v, cap:%v\n", <-ch, len(ch), cap(ch))
	ch <- 5
	fmt.Printf("Buffered : ch:%v, len:%v, cap:%v\n", <-ch, len(ch), cap(ch))

}

