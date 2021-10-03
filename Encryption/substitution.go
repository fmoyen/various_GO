// Chiffrement par substitution
// Author: Fabrice MOYEN
// 2021/10/03
// clé= tableau de substitution
// Taille de l'espace de clé: 26!=2^88=10^27
//   => Attaque par force brute impossible sans moyens informatiques
//   => Attaque par analyse fréquentielle des lettres

package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// #########################################################################################################################
	// Variables
	var Message string
	var MessageOutput string
	var MessageProvided string
	var MessageCryptedProvided string
	var action int
	var count int

	// #########################################################################################################################
	// Introduction
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Chiffrement par substitution")
	fmt.Println("La clé est le tableau de substitution")

	// #########################################################################################################################
	// Crypting map = encryption key
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	Key := map[string]string {
		"a" : "r",
		"b" : "f",
		"c" : "x",
		"d" : "p",
		"e" : "c",
		"f" : "l",
		"g" : "d",
		"h" : "j",
		"i" : "m",
		"j" : "q",
		"k" : "a",
		"l" : "v",
		"m" : "i",
		"n" : "w",
		"o" : "k",
		"p" : "b",
		"q" : "s",
		"r" : "z",
		"s" : "g",
		"t" : "u",
		"u" : "e",
		"v" : "t",
		"w" : "h",
		"x" : "o",
		"y" : "y",
		"z" : "n",
	}
	fmt.Println("           key: ",Key)

	// #########################################################################################################################
	// Building the reverse map for decrypting operations
	DecryptKey := make(map[string]string,len(Key))
	for k,v := range Key {
		DecryptKey[v] = k
	}
	fmt.Println("Decrypting key: ",DecryptKey)
	fmt.Println()

	// #########################################################################################################################
	// Asking for what to do
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Do you want to:")
	fmt.Println("  1- crypt")
	fmt.Println("  2- decrypt")
	fmt.Print("?: ")
	fmt.Scanln(&action)
	fmt.Println()

	if action!=1 && action!=2 {
		fmt.Println ("Acceptable answer should be 1 or 2. Exiting...")
		os.Exit(1)
	}

	// #########################################################################################################################
	// CRYPTING
	if action==1 {
		MessageProvided = "je suis venu jai vu jai vaincu"
		fmt.Println("Example of message: ",MessageProvided)
		fmt.Println("Warning: no uppercase, just letters from 'a' to 'z' or white space ' '")
		fmt.Println()
		fmt.Println("Which message do you want me to crypt:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			MessageProvided = scanner.Text()
		}
		//fmt.Scanln(&MessageProvided) // fmt functions does not support white spaces on purpose

		Message = ""
		for _,c := range MessageProvided {
			if string(c) != " " {
				Message = Message + Key[string(c)]
			}
		}

		count = 0
		MessageOutput = ""
		for _,c := range Message {
			count = count + 1
			MessageOutput = MessageOutput + string(c)
			if count == 5 {
				count = 0
				MessageOutput = MessageOutput + " "
			}
		}

		fmt.Println()
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("Clear Message Provided:   ", MessageProvided, " --> Encrypted Message: ", Message)
		fmt.Println("Clear Message Provided:   ", MessageProvided, " --> Encrypted Message Output: ", MessageOutput)

	// #########################################################################################################################
	// DECRYPTING
	} else if action==2 {
		fmt.Println("Give me the crypted message you want me to decrypt:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			MessageCryptedProvided = scanner.Text()
		}

		Message = ""
		for _,c := range MessageCryptedProvided {
			if string(c) != " " {
				Message = Message + DecryptKey[string(c)]
			}
		}

		fmt.Println()
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("Crypted Message Provided: ", MessageCryptedProvided, " --> Decrypted Message: ", Message)

	}

	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}
