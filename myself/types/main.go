package main

import (
	"fmt"
	"math/rand"
)

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

	months := []Month{
		Jan,
		Abr,
		Sep,
	}

	for _, month := range months {
		fmt.Printf("%v", month.Name())
	}

}

// functional types
// declaration type <Name> func() <Type>
type Generator func() int

// Counter the following function creates a counter
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

//Function extension > the following code implements a code receptor to represent a PH sample
type PH float32

func (p PH) PhCategory() string {
	switch {
	case p < 7:
		return "acid"
	case p > 7:
		return "basic"
	default:
		return "neutral"
	}
}

func PrintPHs() {
	samples := []PH{
		PH(7), PH(9), PH(1.2),
	}

	for _, sample := range samples {
		fmt.Printf("Sample > %v is %v  ", sample, sample.PhCategory())
	}

}

//Pseudoenumerations are used to enumerate values
type Month int

//By using iota the enumeration will do through all the entries
const (
	Jan Month = iota + 1
	Feb
	Mar
	Abr
	May
	Jun
	Jul
	Ago
	Sep
	Oct
	Nov
	Dec
)

func (month Month) Name() string {
	switch month {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	default:
		return "Unknown"
	}
}
