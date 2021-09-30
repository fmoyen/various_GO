package main

import (
	"fmt"
	"os"
)

func main() {
	// Args global variable provided by "os" package holds the command-line arguments, starting with the program name (var Args []string)
	// ./bin/osArgs -a -b
	// os.Args =  [./bin/osArgs -a -b]
	// Length os.Args =  3

	fmt.Println("os.Args = ", os.Args)
	fmt.Println("Length os.Args = ", len(os.Args))
}

