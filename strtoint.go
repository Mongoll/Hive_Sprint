package sprint

/* import (
	"fmt"
) */

func StrToInt(s string) int {
	num := 0

	sign := 1
	start := 0

	// handle optional sign
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {	//loop acording string length
		if s[i] < '0' || s[i] > '9' { // check If the string is not a valid number
			return 0
		}	
		digit := s[i] - '0'        // convert character to integer
		num = num*10 + int(digit)  // shift previous digits left and add new
	}

	return num
}

/* func main() {
	fmt.Println(StrToInt("10203"))
} */