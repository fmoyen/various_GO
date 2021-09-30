package main

import "fmt"

func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	CardMap := make(map[string]string)

	// Set key/value pairs using typical name[key] = val syntax.
	CardMap["U200_capi2"]    = "0x0665"
	CardMap["U50_capi2"]     = "0x0669"
	CardMap["AD9V3_capi2"]   = "0x060f"
	CardMap["AD9H3_capi2"]   = "0x0667"
	CardMap["AD9H7_capi2"]   = "0x0668"
	CardMap["AD9V3_ocapi"]   = "0x060f"
	CardMap["AD9H3_ocapi"]   = "0x0667"
	CardMap["AD9H7_ocapi"]   = "0x0666"
	CardMap["N250SOC_ocapi"] = "0x0666"

	//Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("CardMap:", CardMap)

	// Get a value for a key with name[key].
	v1 := CardMap["AD9V3_ocapi"]
	fmt.Println("9V3 OpenCAPI: ", v1)

	// Get a value for a key with name[key].
	v2 := "AD9H7_ocapi"
	fmt.Println("AD9H7_ocapi value (0x666): ", CardMap[v2])

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(CardMap))

	// The optional second return value when getting a value from a map indicates if the key was present in the map. 
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs := CardMap["k2"]
	fmt.Println("BAD prs:", prs)
	_, prs = CardMap["AD9H3_capi2"]
	fmt.Println("GOOD prs:", prs)
}
