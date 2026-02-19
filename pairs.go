//for test in sandbox and code review I chenged package to "main" and add "funk main{}"
package main

import (
	"fmt"
)

func Pairs() string {
	var result string								//We use "var" because we use "result" outside the loop.
	//result := "" - at that case the same as var
	for i := 0; i <= 98; i++ {						//The outer loop is the first number in the pair.
		for j := i + 1; j <= 99; j++ {				//The inner loop is the second number, which is always greater than the first, so that the pairs are in ascending order.
			pair := fmt.Sprintf("%02d %02d", i, j)	//"Sprintf" format and returns numbers as a string 
			if result == "" {						//| %d → decimal integer | 02 → minimum width 2, adds a leading zero if the number is less than 10
				result = pair						// if "result" empty → make it as first pair
			} else {
				result += ", " + pair				//if "result" not empty → add ", " then add "pair"
			}
		}
	}

	return result
}

func main() {
	fmt.Println(Pairs())
}