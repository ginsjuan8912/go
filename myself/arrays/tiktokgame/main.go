package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var isFirstMove = true
var iaNextMoves []string
var iaLastMove int
var finalMoves map[string]string

/*
The following code emulates a tik tak toe game using arrays, slices, manipulation of string,
and basic usage of maps.2
*/

func main() {
	//the board game is a matrix 3x3
	boardGame := [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	fmt.Println("		¡TIK TOK GAME!")

	//Computer AI possibles plays
	corners := possibleCorners()
	finalMoves = generateFinalGame(&corners)
	//Declare player sign
	player := "X"
	computerPlay := "O"

	fmt.Println(`
                          0,0 | 0,1 | 0,2
                          1,0 | 1,1 | 1,2
                          2,0 | 2,1 | 2,2 `)

	println("Computer Plays as:", computerPlay)
	println("Player plays as:", player)

	play := 10

	for play == 10 {

		computerPlays(&corners, &boardGame)
		printBoard(&boardGame)
		play = validatePlay(&boardGame)

		if DoesPlayerWins(play) {
			break
		}

		if play == 0 {
			fmt.Println("Game Over!")
			break
		}

		playerPlays(&boardGame)
		printBoard(&boardGame)
		play = validatePlay(&boardGame)

		if DoesPlayerWins(play) {
			break
		}

		if play == 0 {
			fmt.Println("Game Over!")
			break
		}
	}

}

func DoesPlayerWins(play int) bool {
	if play == -1 {
		fmt.Println("Computer WINS!")
		return true
	}

	if play == 1 {
		fmt.Println("You WIN!")
		return true
	}

	return false
}

func printBoard(board *[3][3]int) {
	//print each row
	fmt.Println("")
	fmt.Println("")

	row := ""

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			number := board[i][j]
			empty := "-"
			computer := "O"
			player := "X"

			switch number {
			case 0:
				row += empty + "|"
			case 1:
				row += player + "|"
			case -1:
				row += computer + "|"
			}
		}

		fmt.Println(row)
		row = ""
	}

	fmt.Println("")
}

func playerPlays(board *[3][3]int) {

	var x, y int

	printSleepMessage()

	playerPlayed := false

	for playerPlayed == false {
		fmt.Println("")
		fmt.Println("Is your turn, provide an x, y coordinate")
		fmt.Println(`
                          0,0 | 0,1 | 0,2
                          1,0 | 1,1 | 1,2
                          2,0 | 2,1 | 2,2 `)

		fmt.Println("")

		fmt.Println("x:")
		fmt.Scanf("%d\n", &x)
		fmt.Println("y:")
		fmt.Scanf("%d\n", &y)

		fmt.Printf("Player plays on: %v, %v", x, y)

		if board[x][y] == -1 {
			fmt.Println("This box is already taken by computer")
			continue
		}

		board[x][y] = 1
		playerPlayed = true
	}

}

func computerPlays(corners *[]string, board *[3][3]int) {

	numbersReady := false
	fmt.Println("")

	if isFirstMove {

		for numbersReady == false {

			x, y := getNextMove(*corners)

			if board[x][y] == 0 {
				numbersReady = true
				isFirstMove = false

				fmt.Printf("Computer starts on: %v, %v", x, y)
				board[x][y] = -1
			}
		}
	} else {
		//Get the next move from the corners array

		validMove := false
		for validMove == false {
			x, y := getNextMove(*corners)
			if board[x][y] == 0 {
				printSleepMessage()

				fmt.Printf("Computer plays on: %v, %v", x, y)
				board[x][y] = -1

				validMove = true
			}
		}

	}

}

func generateNumber(min int, max int) int {

	rand.Seed(time.Now().UnixMilli())
	return rand.Intn(max-min+1) + min
}

func possibleCorners() []string {
	corners := make([]string, 0, 8)

	corners = append(corners, "0,0;0,2;2,2")
	corners = append(corners, "0,0;2,0;2,2")
	corners = append(corners, "2,0;2,2;0,0")
	corners = append(corners, "2,0;0,0;2,2")
	corners = append(corners, "2,2;0,2;0,0")
	corners = append(corners, "2,2;2,0;0,2")
	corners = append(corners, "0,2;0,0;2,0")
	corners = append(corners, "0,2;0,0;2,0")

	return corners
}

func getNextMove(corners []string) (x int, y int) {

	if isFirstMove {
		iaNextMoves = nil
		index := generateNumber(0, 7)

		corner := corners[index]

		move := corner[0:3]

		//get next corners
		iaNextMoves = strings.Split(corner, ";")
		iaLastMove = 0

		//get final moves
		iaNextMoves = append(iaNextMoves, strings.Split(finalMoves[move], ";")...)

		return getXYValue(move)
	}

	iaLastMove += 1

	return getXYValue(iaNextMoves[iaLastMove])

}

func getXYValue(corner string) (int, int) {
	xy := strings.Split(corner, ",")
	x, err := strconv.Atoi(xy[0])
	y, err2 := strconv.Atoi(xy[1])

	if err != nil && err2 != nil {
		return 0, 0
	}

	return x, y
}

func validatePlay(board *[3][3]int) int {

	play := 10

	//if play == 0 no winner, if play > 1 player wins, and if play -1 computer wins, if play equals 10 then continue game
	var sum = 0

	//r to l validation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += board[i][j]
		}

		if sum == 3 {
			return 1
		}

		if sum == -3 {
			return -1
		}

		sum = 0
	}

	//u to b validation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += board[j][i]
		}

		if sum == 3 {
			return 1
		}

		if sum == -3 {
			return -1
		}

		sum = 0
	}

	//u to b validation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += board[j][i]
		}

		if sum == 3 {
			return 1
		}

		if sum == -3 {
			return -1
		}

		sum = 0
	}

	//perpendicular validation

	sum += board[0][0] + board[1][1] + board[2][2]

	if sum == 3 {
		return 1
	}

	if sum == -3 {
		return -1
	}

	sum = 0

	sum += board[2][0] + board[1][1] + board[0][2]

	if sum == 3 {
		return 1
	}

	if sum == -3 {
		return -1
	}

	continueGame := false

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				continueGame = true
			}
		}
	}

	if !continueGame {
		return 0
	}

	return play

}

func generateFinalGame(corners *[]string) map[string]string {
	nextGame := make(map[string]string)

	for _, corner := range *corners {
		if strings.HasPrefix(corner, "0,0") {
			nextGame[corner[0:3]] = "0,1;1,0;1,1"
		}

		if strings.HasPrefix(corner, "0,2") {
			nextGame[corner[0:3]] = "1,2;0,1;1,1"
		}

		if strings.HasPrefix(corner, "2,0") {
			nextGame[corner[0:3]] = "2,1;1,0;1,1"
		}

		if strings.HasPrefix(corner, "2,2") {
			nextGame[corner[0:3]] = "1,2;2,1;1,1"
		}
	}

	return nextGame
}

func printSleepMessage() {
	fmt.Println("")
	fmt.Println("...")
	time.Sleep(2 * time.Second)
	fmt.Println("")
}
