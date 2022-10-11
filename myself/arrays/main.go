package main

import (
	"fmt"
)

func main() {

	//Declare an array
	anArray := [...]int{1, 2, 3}
	//Declare an empty slice
	var anSlice []int

	//Append is used to add new elements to an slice
	anSlice = append(anSlice, 1, 2, 3)

	//Make will create an slice with a fixed capacity of 2048 and length of 0
	fixedSlice := make([]int, 0, 2048)

	//Vectors are passed as value types to a function, meanwhile, Slices are passed as reference

	fmt.Printf("\nthe array has the following length %v", len(anArray))
	fmt.Printf("\nThe slice has the following length %v, and the following capacity %v", len(anSlice), cap(anSlice))
	fmt.Printf("\nThe slice with Make has the following length %v, and the following capacity %v", len(fixedSlice), cap(fixedSlice))

	Zero(anArray, anSlice)

	//To create a view use the syntax slice := baseSlice[1:3]
	//baseSlice[2:] from index 2 to end
	//baseSlice[:2] from start index to index 2
	//baseSlice[:] from start to end
	view := anSlice[:2]

	fmt.Println("the view: ", view)

	//To make an interation use a for with
	// for <index>, <value> := range <slice> {
	//}

	for _, number := range anSlice {
		fmt.Printf("\n the number is %v", number)
	}

	//Calling a function with a range of values accepted
	add(1, 2, 3, 4)

	//When calling a function that accepts ...[type] passing a array it should be like func(arrayName...)
	add(anSlice...)

}

func Zero(arr [3]int, slice []int) {
	println("\nVectors are passed as value types to a function, meanwhile, Slices are passed as reference")

	arr[0] = 0
	if len(slice) > 0 {
		slice[0] = 0
	}

	fmt.Println("array", arr)
	fmt.Println("slice", slice)

}

/*
 When declaring a variable with a range of values you should use ...<type>
*/
func add(numbers ...int) {

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Println("\nthe sum was", sum)
}
