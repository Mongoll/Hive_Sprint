//for test in sandbox and code review I chenged package to "main" and add "funk main{}"
package main

import (
	"fmt"
) 

func Combinations() string {
	var result string										//We use "var" because we use "result" outside the loop.
	//result := "" - at that case the same as var
	for i := 0; i <= 7; i++ { 								//The loop is the first number in the triplet.
		for j := i + 1; j <= 8; j++ { 						//The loop is the second number in the triplet.
			for k := j + 1; k <= 9; k++ { 					//The loop is the therd number in the triplet.
				triplet := fmt.Sprintf("%d%d%d", i, j, k)	//"Sprintf" format and returns numbers as a string
															// "%d%d%d" - placeholder for decimal integer
				if result == "" {
					result = triplet						// if "result" empty → make it as first triplet.
				} else {
					result += ", " + triplet				//if "result" not empty → add ", " then add "triplet"
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(Combinations())
} 