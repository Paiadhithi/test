package main

import (
	"fmt"

	"time"
)

func main() {
	// creating goroutine from named function
	//create goroutine to execute --> interleaving is seen
	go producer(1)
	go producer(2)

	// Creating goroutines using anonymous functions
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Foo() -message: %v\n", i)

		}
		// another function call
		producer(3)
	}()

	// give groutines time to complete work
	time.Sleep(1 * time.Second)

}

func producer(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Producer # %v - message # %v\n", id, i)
	}
}

