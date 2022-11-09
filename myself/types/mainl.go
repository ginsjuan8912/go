package types

import (
	"fmt"
	"math/rand"
)

// functional types
// declaration type <Name> func() <Type>
type Generator func() int

/*
The following code shows how custom types are declared
*/
func main() {

	//type <Name> <type>
	type Year int
	// Based on a type we could a use portion
	type YearCollection []Year

	//Based on a type
	years := YearCollection{1998, 1992, 2002}
	years = append(years, 2001)
	years = append(years, 2003)

	//When dealing with  custom type conversion have to be specific btw types

	sum := Year(1)

	//The following for sums all the elements from the years instance
	for _, number := range years {
		sum += number
	}

	//This process could be used to create also a type based on a map
	type CityPopulation map[string]int

	population := CityPopulation{
		"Bogota":       11000,
		"Lima":         2500,
		"Buenos Aires": 5800,
	}

	for cityName, pop := range population {
		fmt.Printf("The city %s has %d inhabitants", cityName, pop)
	}

	//the following function creates a counter

}

func Counter() Generator {
	count := 0

	//return an anonymous function
	return func() int {
		count++
		return count
	}
}

func Random(seed int64) Generator {
	rnd := rand.NewSource(seed)
	return func() int {

		return int(rnd.Int63())
	}
}
