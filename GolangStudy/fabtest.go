package main

import (
	"fmt"

)

func main() {

	//================================================================
	pciID := "0003:01:00.0"
	fmt.Println("pciID= ",pciID)
	fmt.Println("len(pciID)= ",len(pciID))
	DBD := pciID[:len(pciID)-2]
	fmt.Println("DBD= ",DBD)


}

