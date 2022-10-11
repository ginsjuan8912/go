package main

import (
	fmt "fmt"
	"math/rand"
	"time"
)

func main() {

	welcome := "Welcome to GuessNumber Game"
	pad := "*******************************"

	fmt.Println(welcome)
	fmt.Println(pad)

	guessNumber := generateNumber()

	fmt.Println("Please guess a number btw 1 and 10, you have three attemps!")

	var userInput int
	numberOfAttemps := 1

	for {

		player := playerPlays(userInput, &guessNumber)

		if player {
			break
		}

		computer := computerPlays(&guessNumber)

		if computer {
			break
		}

		numberOfAttemps++

		if numberOfAttemps > 3 {
			fmt.Println("Oh no! number attemps are over!")
			fmt.Println("GAME OVER")
			break
		}

	}

}

func generateNumber() int {

	rand.Seed(time.Now().UnixMilli())
	min := 1
	max := 10
	return rand.Intn(max-min+1) + min
}

func playerPlays(userInput int, guessNumber *int) bool {
	fmt.Printf("your turn:")
	fmt.Scanf("%d\n", &userInput)

	printSleepMessage()

	if userInput == *guessNumber {
		fmt.Println("That is correct!!! you WIN!!!")
		return true
	}

	return false
}

func computerPlays(guessNumber *int) bool {

	fmt.Println("Is computer turn:")
	computerGuess := generateNumber()

	printSleepMessage()

	fmt.Printf("The computer says: %d", computerGuess)

	printSleepMessage()

	if computerGuess == *guessNumber {
		fmt.Println("That is correct!!! computer WIN!!!")
		return true
	}

	return false
}

func printSleepMessage() {
	fmt.Println("")
	fmt.Println("...")
	time.Sleep(2 * time.Second)
	fmt.Println("")
}
