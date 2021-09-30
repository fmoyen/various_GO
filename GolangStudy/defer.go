package main
import "fmt"

func sample(i string) {  
    fmt.Println(i)
    fmt.Println("Inside the sample()")
}

func main() {  
    //sample() will be invoked only after executing the statements of main()
    myVar := "World"
    defer sample(myVar)
    fmt.Println("Inside the main()")
    myVar="Hello"
    fmt.Println(myVar)
}
