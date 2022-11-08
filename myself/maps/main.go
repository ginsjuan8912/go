package main

import "fmt"

func main() {
	//This code show the use of dictionaries
	//declarartion
	newMap := map[int]string{
		0: "Bogota",
		1: "Madrid",
		2: "Mosquera",
	}

	fmt.Print(newMap)

	//Check if key exists by using a key, value := map
	city, exists := newMap[3]

	if exists {
		fmt.Printf("En el indice %s se encuentra", city)
	} else {
		fmt.Printf("No se encuentra el indice %d")
	}

	//delete an entry from the map with delete(mapName, key)
	delete(newMap, 1)

	//iterate through a map with range

	for key, value := range newMap {
		fmt.Println("")
		fmt.Printf("the key is %d and the value is %s", key, value)
	}

	//to count the number of elements in the map use the function len
	numberOfElements := len(newMap)

	fmt.Printf("\nThe number of elements is %d", numberOfElements)

	//create a map with a specific length with make
	createdMap := make(map[string]string, 10000)
	fmt.Printf("\nThe following map has a length of %d", len(createdMap))

}
