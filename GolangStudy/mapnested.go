package main

import "fmt"

func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]map[string]int)
	// You have to initialize the maps stored inside the map in order to store data in them (2 syntaxes to do so)
	m["A"] = map[string]int{}
	m["B"] = make(map[string]int)

	// Set key/value pairs using typical name[key][key] = val syntax.
	m["A"]["a"] = 7
	m["A"]["b"] = 2
	m["A"]["c"] = 9
	m["B"]["a"] = 13
	m["B"]["b"] = 5

	// In the case that the map is being populated iteratively, you can check for a nil value prior to assignment
	if m["C"] == nil {
		m["C"] = map[string]int{}
	}
	m["C"]["a"] = 8
	m["C"]["b"] = 12

	//Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("map:", m)

//===========================================================================================================================

	// To create an empty map, and initialise the second level maps in one command
	var data = map[string]map[string]int{
		"A": map[string]int{},
		"B": map[string]int{},
	}

	// Set key/value pairs using typical name[key][key] = val syntax.
	data["A"]["a"] = 7
	data["A"]["b"] = 2
	data["A"]["c"] = 9
	data["B"]["a"] = 13
	data["B"]["b"] = 5

	//Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("data:", data)
	fmt.Println("data[A][c]:", data["A"]["c"])

//===========================================================================================================================

	// Also possible to use a simple map with a structure as key for the map. So only one map to initialize
	// And also no need to check the initialization of the second level map before using it as there is only one level
	type key struct {
		first, second string
	}

	simplemap := make(map[key]int)

	simplemap[key{"A", "a"}] = 12
	simplemap[key{"A", "b"}] = 8
	simplemap[key{"B", "a"}] = 5

	fmt.Println("Simple Map with key as struct:", simplemap)

}
