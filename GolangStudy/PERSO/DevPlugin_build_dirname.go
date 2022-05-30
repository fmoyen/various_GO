package main

import (
	"fmt"
        "path"
)

func main() {

	const (
		RootDir = "/sys/devices"
		Prefix = "pci"
	)

	//================================================================
	// 0004:00:00.1 -> /sys/devices/pci0004:00 
	ID := "0004:00:00.1"
	fmt.Println("ID= ",ID)
	fmt.Println("len(ID)= ",len(ID))
	Suffix := ID[:len(ID)-5]
	fmt.Println("Suffix= ",Suffix)

	DevPath := path.Join(RootDir, Prefix + Suffix)
	fmt.Printf("Dev Path = %s\n", DevPath)


}

