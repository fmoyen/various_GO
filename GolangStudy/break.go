package main
import "fmt"

func main() {  
	L:
	for i := 1; i <= 10; i++ {
		fmt.Println("loop I --> ",i)
		for j :=1; j <= 10; j++ {
			fmt.Println("    loop J --> ",j)
			if j == 5 {
				break L
			}
		}
	}

	fmt.Println("-------------------")

	for k := 1; k <= 10; k++ {
		fmt.Println("loop K --> ",k)
		if k == 7 {
			break
		}
	}

}
