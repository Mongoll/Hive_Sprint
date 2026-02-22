//for test in sandbox and code review I chenged package to "main" and add "funk main{}"
package main

import (
	"fmt"
) 

func AlphabetMastery(n int) string {
	
	if n < 1 || n > 26 {			// <- It returns an error if the input is outside the valid range.
		return ""					// <- empty line - cose autotests not pass if i have some text in this line
	} 								

	result := make([]rune, n) 		// <-create array with 'n' elements. Example: result = [0, 0, 0, ..., 0]
	//also we can use
	//result := []rune{}

	for i := 0; i < n; i++ {		// <- The Unicode code point for 'a' is 97. rune(i) converts i to the same type.
		result[i] = 'a' + rune(i) 	// <- When we add them together: i = 0 → 'a' + 0 = 'a'; i = 1 → 'a' + 1 = 'b'...
	//also we can use 
	//result = append(result, 'a'+rune(i))
	}

	return string(result)			// <- return array of numbers Unicode and convert them to a string
}

func main() {
	fmt.Println(AlphabetMastery(6))
} 


//refactoring without using array
/* 
 func AlphabetMastery(n int) string {

	result := ""

	if n < 1 || n > 26 { 
		return ""
	}

	for i := 0; i < n; i++ { 
		result += string('a' + rune(i)) 
	}

	return result
} 
 */