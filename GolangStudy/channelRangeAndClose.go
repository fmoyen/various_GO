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
	close(c)
}

func main() {
	c := make(chan int, 10)  // create a channel of int with a buffer of 10 elements
	fmt.Println("capacity of c channel:",cap(c))

	go fibonacci(cap(c), c)

	// The loop for i := range c receives values from the channel repeatedly until it is closed. 
	for i := range c {
		fmt.Println(i)
	}
}

/* 
Channels aren't like files; you don't usually need to close them.
Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop. 
*/
