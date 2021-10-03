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
)

func main() {
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Chiffrement par substitution")
	fmt.Println("La clé est le tableau de substitution")
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

	// Building the reverse map for decrypting operations
	DecryptKey := make(map[string]string,len(Key))
	for k,v := range Key {
		DecryptKey[v] = k
	}
	fmt.Println("Decrypting key: ",DecryptKey)
	fmt.Println()


	MessageProvided := "jesuisvenujaivujaivaincu"
	var MessageCryptedProvided string
	fmt.Println("Example of message: ",MessageProvided)
	fmt.Println("Warning: no space, no uppercase, just letters from 'a' to 'z'")
	fmt.Println()

	var Message string
	var action int

	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Do you want to:")
	fmt.Println("  1- crypt")
	fmt.Println("  2- decrypt")
	fmt.Print("?: ")
	fmt.Scanln(&action)

	if action!=1 && action!=2 {
		fmt.Println ("Acceptable answer should be 1 or 2. Exiting...")
		os.Exit(1)
	}

	// CRYPTING
	if action==1 {
		fmt.Println("Which message do you want me to crypt:")
		fmt.Scanln(&MessageProvided)
		fmt.Println("------------------------------------------------------------------------------------------------------------")

		Message = ""
		for _,c := range MessageProvided {
			Message = Message + Key[string(c)]
		}

		fmt.Println("Clear Message Provided:   ", MessageProvided, " --> Encrypted Message: ", Message)
		fmt.Println()

	// DECRYPTING
	} else if action==2 {
		fmt.Println("Give me the crypted message you want me to decrypt:")
		fmt.Scanln(&MessageCryptedProvided)
		fmt.Println("------------------------------------------------------------------------------------------------------------")

		Message = ""
		for _,c := range MessageCryptedProvided {
			Message = Message + DecryptKey[string(c)]
		}

		fmt.Println("Crypted Message Provided: ", MessageCryptedProvided, " --> Decrypted Message: ", Message)
		fmt.Println()

	}

	fmt.Println()
}
