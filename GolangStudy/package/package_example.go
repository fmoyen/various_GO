package main
import "fmt"
//the package to be created
import "calculation"

func main() {  
	x,y := 15,10
	//the package will have function Do_add()
sum := calculation.Do_add(x,y)
fmt.Println("Sum",sum) 
fmt.Println("calculation.Calculs",calculation.Calculs) // it's possible de define global variables in the called package and use them in the main()
}
