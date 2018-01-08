package main

import (
	"fmt"
	"math"
)

func bs(numbers []int, n, min, max int) (int, error) {

	if n > max {
		return 0, fmt.Errorf("not in the array")
	}

	guessPosition := avg(min, max)

	if numbers[guessPosition] == n {
		return guessPosition, nil
	}

	if numbers[guessPosition] > n {
		max = guessPosition - 1
		result, err := bs(numbers, n, min, max)
		return result, err
	} else {
		min = guessPosition + 1
		result, err := bs(numbers, n, min, max)
		return result, err
	}

}

func avg(min, max int) int {

	var m int

	if (min+max)%2 == 0 {
		m = (min + max) / 2
	} else {
		m = int(math.Floor(float64(min+max) / 2.0))
	}

	return m
}
