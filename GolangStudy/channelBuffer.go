package main

import "fmt"

func main() {
	ch := make(chan int, 2)  // The channel is created with a buffer of 2 int elements
	// So that's possible to push a second element into the buffer while the first pushed element is still in the channel (hasn't been read yet)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
If not using a channel with a buffer of 2 elements, running the program will fail with a deadlock fatal error
(it will never be possible to push the second element while there is nothing in the program capable of reading the first element from the channel and so emptying the channel)
*/
