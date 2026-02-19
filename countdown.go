package sprint

/* import (
	"fmt"
) */

func Countdown(n int) string {
	if n < 0 || n > 9 {
		return "Input must be a one-digit number"
	}

	var result string
	//result := ""  - at that case the same as var
	for i := n; i > 0; i -= 2 { 				//"i -= 2" the same as "i = i - 2"
		if result != "" {
			result += ", "						//if "result" contain data → add ", " at the end
		}
		result += string('0' + i)
		//result += fmt.Sprintf("%d", i)			//"Sprintf" format and returns numbers as a string | "%d" - placeholder for decimal integer
	}

	// add zero at the end of line
	if result != "" {
		result += ", 0!"						//if "result" contain data → add ", " and zero at the end
	} else {
		result = "0!"							//if "result" empty → add zero without ", " at the end
	}

	return result
}
/* 
func main() {
	fmt.Println(Countdown(8)) // Output: 8, 6, 4, 2, 0!
	fmt.Println(Countdown(5)) // Output: 5, 3, 1, 0!
	fmt.Println(Countdown(0)) // Output: 0!
} */
