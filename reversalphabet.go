package sprint

/* import (
	"fmt"
) */

func ReverseAlphabet(step int) string {

	if step <= 0 {									//If step is 0 or negative, use a default value of 1
		step = 1
	}									
													//Create an empty list (array) for processing/fetching letters
	var result []rune								//We use "var" because we use "result" outside the loop.
	//result := []rune{} - at that case the same as var
	
	for ch := 'z'; ch >= 'a'; ch -= rune(step) {  	//For "ch", we don’t use "var" because "ch" exists only inside the loop.
		result = append(result, ch)					// ch -= rune(step) the same as ch = ch - rune(step)
	}												// The append built-in function appends elements to the end of a slice. in that case: "result" - slise, "ch" - element

	return string(result)							//Convert an array of Unicode numbers to string letters
}

/* func main() {
	fmt.Println(ReverseAlphabet(5))
} */