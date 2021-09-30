package main

import (
	"fmt"
	"strings"
	"path"
)

func main() {

const (
	CXLDevDir      = "/dev/cxl"
	CXLPrefix1     = "afu"
	CXLPrefix2     = ".0m"
      )

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Printf("%q\n", strings.Split("card0", "card"))

	fmt.Printf("%q\n", strings.TrimPrefix("card0", "card"))

	capiIDSTR := "1"
	CXLDevPath := path.Join(CXLDevDir, CXLPrefix1 + capiIDSTR + CXLPrefix2)
	fmt.Printf("%s\n", CXLDevPath)
}

