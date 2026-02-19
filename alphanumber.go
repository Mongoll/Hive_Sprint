package sprint

/* import (
	"fmt"
) */

func AlphaNumber(n int) string {

	if n == 0 {							//check if 0 put "a" to string
		return "a"						//if skip this check, digits is empty, 
	}									//and the final string(digits) would return "" (empty string) instead of "a"

	negative := n < 0					//need a positive number for the modulo/division loop 
	if negative {						//(n % 10, n /= 10), because % with negative numbers is tricky in Go.
		n = -n
	}

	digits := make([]byte, 0, 20)		//create array for separate number into digits

	
	for n > 0 {							// separate number into digits
		digit := n % 10
		digits = append(digits, byte('a'+digit))
		n /= 10
	}

	// reverse digits needed because when extract digits using n % 10, you get them in reverse order (from least significant to most significant).
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	if negative {						//still need to flip n to positive to extract digits correctly
		return "-" + string(digits)
	}
	return string(digits)
}

/* func main(){
	fmt.Println(AlphaNumber(234))
} */