// Chiffre de Vigenère
// Author: Fabrice MOYEN
// 2021/10/03
// Chiffrage par substitution où chaque lettre du message est codé grace à la lettre correspondante de la clé
//   - Le chiffrement consiste à additionner chaque lettre du message avec la lettre de la clé en dessous, modulo 26
//   - Le déchiffrement consiste à soustraire chaque lettre du message chiffré avec la lettre de la clé en dessous, modulo 26
//
//   => Attaque par force brute impossible sans moyens informatiques
//   => Attaque par analyse fréquentielle des lettres facilitée si on connait la longueur de la clé

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
	var MessageLength int
	var KeyProvided string
	var KeyString string
	var KeyStringLength int
	var action int

	// #########################################################################################################################
	// Introduction
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Chiffre de Vigenère")
	fmt.Println("La clé est mise en dessous du message et chaque lettre du message est codée avec la lettre correspondante de la clé")

	// #########################################################################################################################
	// Asking for what to do
	fmt.Println()
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

	// #########################################################################################################################
	// Asking for the key
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Key to use (with space chars if you want)? :")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		KeyProvided = scanner.Text()
	}

	//strings.ReplaceAll(randomString, " ", "") ??

	KeyString = ""
	for _,c := range KeyProvided {
		if string(c) != " " {
			KeyString = KeyString + string(c)
		}
	}
	KeyStringLength = len(KeyString)
	fmt.Println()
	fmt.Println("Used formatted key will be: ",KeyString)
	fmt.Println("Length: ",KeyStringLength)

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
		MessageLength = len(MessageProvided)

		//MessageSlice := make([]string, MessageLength)
		MessageSlice := make([]rune, MessageLength)
		for i,c := range MessageProvided {
			//MessageSlice[i] = string(c)
			MessageSlice[i] = c
		}
		fmt.Println("Clear Message Provided Slice: ", MessageSlice)

		KeySlice := make([]string, MessageLength)
		for i:=0; i<MessageLength; i++ {
			KeySlice[i] = string(KeyString[i % KeyStringLength])
		}
		fmt.Println("Key Slice                   : ", KeySlice)

		//CryptedSlice := make([]string, MessageLength)
		//for i:=0; i<MessageLength; i++ {
			//CryptedSlice[i] = string(MessageSlice[i] + KeySlice[i])
		//}


		fmt.Println()
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("Clear Message Provided:   ", MessageProvided, " --> Encrypted Message: ", Message)
		fmt.Println("Clear Message Provided:   ", MessageProvided, " --> Encrypted Formatted Output: ", MessageOutput)

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
