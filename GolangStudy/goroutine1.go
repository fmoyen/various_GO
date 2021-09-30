package main
import "fmt"
import "time"
    
func display() {
	for i:=0; i<5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("In display")
	}
}

func main() {
	//invoking the goroutine display()
	go display()
	//The main() continues without waiting for display()
	for i:=0; i<5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("In main")
	}
}
