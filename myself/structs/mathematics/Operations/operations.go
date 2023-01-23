package Operations

import (
	"errors"
)

func Divide(divisor int, dividend int) (int, error) {

	if divisor == 0 {
		//Throw error when divisor is zero
		return 0, errors.New("Divisor cannot be zero")
	}

	return dividend / divisor, nil
}

func Sum(numbersCh <-chan int) int {
	total := 0

	for number := range numbersCh {
		total += number
	}

	return total
}
