package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// Open a file for writing
	file, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write 100,000 random numbers to the file
	for i := 0; i < 100000; i++ {
		fmt.Fprintf(file, "%d,", rand.Intn(100000))
	}
}
