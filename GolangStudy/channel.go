package main
import "fmt"
import "time"
    
func display(ch chan int) {
	time.Sleep(5 * time.Second)
	fmt.Println("Inside display()")
	ch <- 1234
}

func main() {
	ch := make(chan int) 
	go display(ch)
	x := <-ch  // this command is complete when something has been received from the channal ch. So main is waiting here for goroutine display to finish
	fmt.Println("Inside main()")
	fmt.Println("Printing x in main() after taking from channel:",x)
}

