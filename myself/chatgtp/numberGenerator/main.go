package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	// Open a file for writing
	file, err := os.Create("numbers2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write 100,000 random numbers to the file
	for i := 0; i < 10; i++ {
		fmt.Fprintf(file, "%d,", generateNumber(1, 20))
	}
}

func generateNumber(min int, max int) int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixMilli())
	return rand.Intn(max-min+1) + min
}
