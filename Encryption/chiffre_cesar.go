package main
import (
	"fmt"
	"os"
)

func main() {
	Alphabet := [...] string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var AlphabetCrypted [26] string
	var AlphabetDecrypted [26] string
	MessageProvided := "jesuisvenujaivujaivaincu"
	MessageCryptedProvided := "rmacqadmvcriqdcriqdiqvkc"
	fmt.Println("Exemple of messages: ",MessageProvided," --(key=8)--> ", MessageCryptedProvided)
	fmt.Println("Warning: no space, no uppercase, just letters from 'a' to 'z'")
	var Message string
	var action int
	var key int

	fmt.Println()
	fmt.Println("Do you want to:")
	fmt.Println("  1- crypt")
	fmt.Println("  2- decrypt")
	fmt.Println("  3- attack a crypted message")
	fmt.Println("  4- get just an example")
	fmt.Print("?: ")
	fmt.Scanln(&action)

	if action!=1 && action!=2 && action!=3 && action!=4 {
		fmt.Println ("Acceptable answer should be 1, 2, 3 or 4. Exiting...")
		os.Exit(1)
	}

	// CRYPTING
	if action==1 {
		fmt.Println("Which message do you want me to crypt:")
		fmt.Scanln(&MessageProvided)
		fmt.Println("Give me the key (number between 1 to 25::")
		fmt.Scanln(&key)
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("key: ",key)
		for i :=0; i<=25; i++ {
			increment := i + key
			increment = increment % 26 //modulo 26, so if increment=26, then increment=0, etc
			AlphabetCrypted[i] = Alphabet[increment]
		}
		fmt.Println("Encryption: ", Alphabet, " -> ", AlphabetCrypted)

		Message = ""
		for _,c := range MessageProvided {
			charindex := c - 97 // "a"->c=97, "b"->c=98, etc (string(97)="a", string(98)="b", etc)
			Message = Message + AlphabetCrypted[charindex]
		}

	fmt.Println("Clear Message Provided:   ", MessageProvided, " --> Encrypted Message: ", Message)

	// DECRYPTING
	} else if action==2 {
		fmt.Println("Give me the crypted message you want me to decrypt:")
		fmt.Scanln(&MessageCryptedProvided)
		fmt.Println("Give me the key (number between 1 to 25::")
		fmt.Scanln(&key)
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("key: ",key)
		for j :=26; j<=51; j++ {
			decrement := j - key
			decrement = decrement % 26 //modulo 26, so if increment=26, then increment=0, etc
			AlphabetDecrypted[j - 26] = Alphabet[decrement]
		}
		fmt.Println("Decryption: ", Alphabet, " -> ", AlphabetDecrypted)
		fmt.Println()

		Message = ""
		for _,c := range MessageCryptedProvided {
			charindex := c - 97 // "a"->c=97, "b"->c=98, etc (string(97)="a", string(98)="b", etc)
			Message = Message + AlphabetDecrypted[charindex]
		}

		fmt.Println("Crypted Message Provided: ", MessageCryptedProvided, " --> Decrypted Message: ", Message)
		fmt.Println()

	// ATTACKING
	} else if action==3 {
		fmt.Println("Give me the crypted message you want me to attack:")
		fmt.Scanln(&MessageCryptedProvided)
		for key = 1; key <= 25; key++ {
			fmt.Println("------------------------------------------------------------------------------------------------------------")
			fmt.Println("key: ",key)

			for j :=26; j<=51; j++ {
				decrement := j - key
				decrement = decrement % 26 //modulo 26, so if increment=26, then increment=0, etc
				AlphabetDecrypted[j - 26] = Alphabet[decrement]
			}
			fmt.Println("Decryption: ", Alphabet, " -> ", AlphabetDecrypted)
			fmt.Println()

			Message = ""
			for _,c := range MessageCryptedProvided {
				charindex := c - 97 // "a"->c=97, "b"->c=98, etc (string(97)="a", string(98)="b", etc)
				Message = Message + AlphabetDecrypted[charindex]
			}

			fmt.Println("Crypted Message Provided: ", MessageCryptedProvided, " --> Decrypted Message: ", Message)
			fmt.Println()
		}

	// EXAMPLE
	} else if action==4 {
		fmt.Println("Please find hereunder just an example:")

		for key = 1; key <= 25; key++ {
			fmt.Println("------------------------------------------------------------------------------------------------------------")
			fmt.Println("key: ",key)
			for i :=0; i<=25; i++ {
				increment := i + key
				increment = increment % 26 //modulo 26, so if increment=26, then increment=0, etc
				AlphabetCrypted[i] = Alphabet[increment]
			}
			fmt.Println("Encryption: ", Alphabet, " -> ", AlphabetCrypted)

			for j :=26; j<=51; j++ {
				decrement := j - key
				decrement = decrement % 26 //modulo 26, so if increment=26, then increment=0, etc
				AlphabetDecrypted[j - 26] = Alphabet[decrement]
			}
			fmt.Println("Decryption: ", Alphabet, " -> ", AlphabetDecrypted)
			fmt.Println()

			Message = ""
			for _,c := range MessageProvided {
				charindex := c - 97 // "a"->c=97, "b"->c=98, etc (string(97)="a", string(98)="b", etc)
				Message = Message + AlphabetCrypted[charindex]
			}

		fmt.Println("Clear Message Provided:   ", MessageProvided, " --> Encrypted Message: ", Message)

			Message = ""
			for _,c := range MessageCryptedProvided {
				charindex := c - 97 // "a"->c=97, "b"->c=98, etc (string(97)="a", string(98)="b", etc)
				Message = Message + AlphabetDecrypted[charindex]
			}

			fmt.Println("Crypted Message Provided: ", MessageCryptedProvided, " --> Decrypted Message: ", Message)
			fmt.Println()
		}
	}

	fmt.Println()
}
