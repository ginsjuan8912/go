package main

import "fmt"

func main() {
	//Declaracion
	helloMessage := "Hello"
	worldMessage := "World"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//printf

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)

	//sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Println(message)
}
