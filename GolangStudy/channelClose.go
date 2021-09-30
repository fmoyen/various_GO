package main
import "fmt"
import "time"

//This subroutine pushes numbers 0 to 9 to the channel and closes the channel
func add_to_channel(ch chan int) {
	fmt.Println("Send data")
	for i:=0; i<10; i++ {
		ch <- i //pushing data to channel
	}
	close(ch) //closing the channel to inform the receiver nothing more will be sent to the channel

}

//This subroutine fetches data from the channel and prints it.
func fetch_from_channel(ch chan int) {
	fmt.Println("Read data")
	for {
		//fetch data from channel (status tells if the channel is opened or closed)
		x, status := <- ch

		//status is true if there is data received from the channel
		//status is false if there is no more data to receive from the channel and the channel is closed
		if status == true {
			fmt.Println(x)
		}else{
			fmt.Println("channel has been closed")
			break
		}
	}
}

func main() {
	//creating a channel variable to transport integer values
	ch := make(chan int)

	//invoking the subroutines to add and fetch from the channel
	//These routines execute simultaneously
	go add_to_channel(ch)
	go fetch_from_channel(ch)  // the receiver can be started after the sender because the sender will no longer send anything in the channel if what has already been sent has not yet been read

	//delay is to prevent the exiting of main() before goroutines finish
	time.Sleep(5 * time.Second)
	fmt.Println("Inside main()")
}
