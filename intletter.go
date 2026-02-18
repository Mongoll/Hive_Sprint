package main

import (
	"errors"
	"fmt"
)

// It returns an error if the input is outside the valid range [1, 26].
func GetLetterFromNumber(n int) (rune, error) {
	if n < 1 || n > 26 {
		return 0, errors.New("input must be a positive integer between 1 and 26")
	}

	// 'a' has an ASCII value of 97.
	// To get the corresponding letter, we add (n - 1) to the ASCII value of 'a'.
	// For example, if n = 1, we get 'a' + 0. If n = 2, we get 'a' + 1 ('b').
	letter := rune('a' + n - 1)

	return letter, nil
}

func main() {
	// Example usage
	inputs := []int{1, 5, 26, 0, 27}

	for _, input := range inputs {
		letter, err := GetLetterFromNumber(input)
		if err != nil {
			fmt.Printf("Input %d: %v\n", input, err)
		} else {
			fmt.Printf("Input %d: %c\n", input, letter)
		}
	}
}