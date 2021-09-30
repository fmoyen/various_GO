package main
import "fmt"

func main() {
	//Create an integer variable a with value 20
	a := 20
	
	//Create a pointer variable b and assigned the address of a
	var b *int = &a

	//print address of a(&a) and value of a  
	fmt.Println("Address of a:",&a)
	fmt.Println("Value of a:",a)

	//print b which contains the memory address of a i.e. &a
	fmt.Println("Address pointed by pointer b (= Address of a):",b)

	//*b prints the value in memory address which b contains i.e. the value of a
	fmt.Println("Value stored at the Address pointed by pointer b:",*b)

	//increment the value of variable a using the variable b
	fmt.Println("incrementing the value pointed by pointer b")
	*b = *b+1

	//prints the new value using a and *b
	fmt.Println("Value stored at the Address pointed by pointer b:",*b)
	fmt.Println("Value of a:",a)}

