package main
import "fmt"

func main() {
	Alphabet := [...] string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var AlphabetCrypted [26] string
	var AlphabetDecrypted [26] string
	MessageProvided := "jesuisvenujaivujaivaincu"
	MessageCryptedProvided := "rmacqadmvcriqdcriqdiqvkc"
	//MessageCryptedProvided := "abcdefghijklmnopqrstuvwxyz"
	var Message string

	fmt.Println()
	for key := 1; key <= 25; key++ {
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
