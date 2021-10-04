/*
##################################################################################################################################
Chiffre de Vigenère

Author: Fabrice MOYEN
Date: 2021/10/03

Chiffrage par substitution où chaque lettre du message est codé grace à la lettre correspondante de la clé
  - Le chiffrement consiste à additionner chaque lettre du message avec la lettre de la clé en dessous, modulo 26
  - Le déchiffrement consiste à soustraire chaque lettre du message chiffré avec la lettre de la clé en dessous, modulo 26

  => Attaque par force brute impossible sans moyens informatiques
  => Attaque par analyse fréquentielle des lettres facilitée si on connait la longueur de la clé

À partir du XIXe siècle, on utilise des machines mécaniques à cylindres de plus en plus complexes basées sur
le principe de substitution polyalphabétique comme le chiffre de Vigenère. La plus connue de ces machines est la machine Enigma.
##################################################################################################################################
*/

package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// #########################################################################################################################
	// Variables
	var MessageProvided string
	var MessageCryptedProvided string
	var MessageLength int
	var KeyProvided string
	var KeyString string
	var KeyStringLength int
	var KeyStringMessageLength string
	var CryptedString string
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

		MessageSlice := make([]string, MessageLength)
		MessageRuneSlice := make([]rune, MessageLength)
		for i,c := range MessageProvided {
			MessageSlice[i] = string(c)
			MessageRuneSlice[i] = c
		}

		KeySlice := make([]string, MessageLength)
		KeyRuneSlice := make([]rune, MessageLength)
		KeyStringMessageLength = ""
		for i:=0; i<MessageLength; i++ {
			//KeySlice[i] = string(KeyString[i % KeyStringLength])
			KeyStringMessageLength = KeyStringMessageLength + string(KeyString[i % KeyStringLength])
		}

		//fmt.Println("Key String Message Length   : ", KeyStringMessageLength)

		for i,c := range KeyStringMessageLength {
			KeySlice[i] = string(c)
			KeyRuneSlice[i] = c
		}

		fmt.Println()
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("Clear Message Provided Slice     : ", MessageSlice)
		fmt.Println("Key Slice                        : ", KeySlice)
		//fmt.Println("------------------------------------------------------------------------------------------------------------")
		//fmt.Println("Clear Rune Message Provided Slice: ", MessageRuneSlice)
		//fmt.Println("Key Rune Slice                   : ", KeyRuneSlice)

		CryptedRuneSlice := make([]rune, MessageLength)
		CryptedSlice := make([]string, MessageLength)
		CryptedString = ""
		for i:=0; i<MessageLength; i++ {
			if MessageSlice[i] != " " {
				CryptedRuneSlice[i] = ( MessageRuneSlice[i] - 97 + KeyRuneSlice[i] - 97 ) % 26 + 97
				/*
				Explanation:
				------------------------------------------------------------------------------------------------------------
				Clear Message Provided Slice     :  [a b c]
				Key Slice                        :  [a b c]
				------------------------------------------------------------------------------------------------------------
				Clear Rune Message Provided Slice:  [97 98 99]   => msgrune-97=[0 1 2]
				Key Rune Slice                   :  [97 99 101]  => keyrune-97=[0 2 4]  (key=a means no change so +0)
				------------------------------------------------------------------------------------------------------------
				Crypted Rune Message Slice       :  [97 99 101]  = (msgrune-97+keyrune-97)%26+97
				Crypted Message Slice            :  [a c e]
				*/
			} else {
				CryptedRuneSlice[i] = MessageRuneSlice[i] // = rune of " "
			}
				CryptedSlice[i] = string(CryptedRuneSlice[i])
				CryptedString = CryptedString + CryptedSlice[i]
		}

		fmt.Println("------------------------------------------------------------------------------------------------------------")
		//fmt.Println("Crypted Rune Message Slice       : ", CryptedRuneSlice)
		fmt.Println("Crypted Message Slice            : ", CryptedSlice)

		fmt.Println()
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("Clear Message Provided:   ", MessageProvided)
		fmt.Println("  --> Crypted Message:    ", CryptedString)

	// #########################################################################################################################
	// DECRYPTING
	} else if action==2 {
		fmt.Println("Give me the crypted message you want me to decrypt:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			MessageCryptedProvided = scanner.Text()
		}
		fmt.Println()
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("Crypted Message Provided:   ", MessageCryptedProvided)

	}

	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}
