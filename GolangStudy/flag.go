package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	fmt.Println("os.Args = ", os.Args)
	fmt.Println("os.Args[1:] = ", os.Args[1:])
	fmt.Println("Length os.Args = ", len(os.Args))

	wordPtr := flag.String("word", "foo", "a string") // Here we declare a string flag word with a default value "foo" and a short description. 

	// This declares numb and fork flags, using a similar approach to the word flag.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// It’s also possible to declare an option that uses an existing var declared elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	// Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program. 
	// This will assign to the flags the command line parameters given (os.Args[1:]) when executing the program 
	flag.Parse()

	// Here we’ll just dump out the parsed options and any trailing positional arguments. Note that we need to dereference the pointers with e.g. *wordPtr to get the actual option values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())  // Args returns the non-flag command-line arguments (so empty is good !)

	/*	
	To experiment with the command-line flags program it’s best to first compile it and then run the resulting binary directly.
	Try out the built program by first giving it values for all flags.
	   ./command-line-flags -word=opt -numb=7 -fork -svar=flag
	Note that if you omit flags they automatically take their default values.

	[fabrice@FabT480 studyGo] GOBIN=$PWD go install flag.go 

	[fabrice@FabT480 studyGo] ./flag 
	os.Args =  [./flag]
	os.Args[1:] =  []
	Length os.Args =  1
	word: foo
	numb: 42
	fork: false
	svar: bar
	tail: []

	[fabrice@FabT480 studyGo] ./flag -word=opt -numb=7 -fork -svar=flag
	os.Args =  [./flag -word=opt -numb=7 -fork -svar=flag]
	os.Args[1:] =  [-word=opt -numb=7 -fork -svar=flag]
	Length os.Args =  5
	word: opt
	numb: 7
	fork: true
	svar: flag
	tail: []

	[fabrice@FabT480 studyGo] ./flag -h
	os.Args =  [./flag -h]
	os.Args[1:] =  [-h]
	Length os.Args =  2
	Usage of ./flag:
  	-fork
    		a bool
  	-numb int
    		an int (default 42)
  	-svar string
    		a string var (default "bar")
  	-word string
    		a string (default "foo")

	Note that the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments).
	[fabrice@FabT480 studyGo] ./flag -word=opt -numb=7 -fork=false a b -svar=flag 
	os.Args =  [./flag -word=opt -numb=7 -fork=false a b -svar=flag]
	os.Args[1:] =  [-word=opt -numb=7 -fork=false a b -svar=flag]
	Length os.Args =  7
	word: opt
	numb: 7
	fork: false
	svar: bar
	tail: [a b -svar=flag]


	*/
}

