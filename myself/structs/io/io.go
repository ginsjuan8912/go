package io

import (
	"errors"
	"fmt"
	"io"
	"os"
)

/*
The following functions writes into a buffer and return the number of bytes
that were written by the function
*/

type IO struct {
	file     *os.File
	Filename string
}

func (f IO) Write(buffer []byte) (n int, err error) {
	//The number of bytes that were written in n
	n, err = os.Stdout.Write(buffer)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Number of bytes written %d", n)

	return 0, nil
}

func (f IO) CreateFile(filename string, contents string) (err error) {
	f.file, err = os.Create(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	f.Filename = filename

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Cannot close the file")
		}
	}(f.file)

	if _, err = f.file.Write([]byte(contents)); err != nil {
		fmt.Println(err)
		return
	}
	return nil
}

func (f IO) Append(contents string) (n int, err error) {

	if f.file, err = os.OpenFile(f.Filename, os.O_APPEND, 0644); err != nil {
		return 0, errors.New(fmt.Sprintf("Cannot append contents to file %s", f.Filename))
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(f.file)

	if n, err = f.file.WriteString(contents); err != nil {
		return 0, err
	}

	return n, nil
}

//ReadFile
func (f IO) ReadFile(filename string) (contents string, err error) {

	contents = ""
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	buffer := make([]byte, 1024)

	for {
		content, err := file.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		if content > 0 {
			contents += string(buffer[:content])
		}
	}

	return contents, err
}
