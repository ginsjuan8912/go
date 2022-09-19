package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	
	content := "Hello from Go!"

	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()
	
	ln, err := io.WriteString(file, content)
	checkError(err)
	
	fmt.Printf("All done with file of %v characters", ln)
	
	bytes := []byte(content)
	ioutil.WriteFile("./fromBytes.txt", bytes, 0644)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}