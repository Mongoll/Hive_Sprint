package main

 import (
	"fmt"
) 

func SimpleStrToInt(s string) int {
num := 0

	for i := 0; i < len(s); i++ {	//loop acording string length
		if s[i] < '0' || s[i] > '9' { // check If the string is not a valid number
			return 0
		}	
		digit := s[i] - '0'        // convert character to integer cose: In Go, characters are stored as numbers (ASCII codes). the same as we have case with alphabet
		num = num*10 + int(digit)  // shift previous digits left and add new
	}

	return num
}
func main() {
	fmt.Println(SimpleStrToInt("-10203"))
} 