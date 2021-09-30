package main
import ("fmt")

func main() {
	a := 20
	fmt.Println(a)

	//gives error since a is already declared
	a := 30
	fmt.Println(a)
}
