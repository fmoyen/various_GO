package main

import (
	"fmt"
)

func main () {

	//var files = [] string {"hosts", "os-level"}
	files := [] string{} {"hosts", "os-level"}

	for i := 1; i <= 5; i++ {
		fmt.Println("loop",i)
	}

	for _, f := range files{
		fmt.Println("fichier",f)
	}

}
