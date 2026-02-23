package main

// import "fmt"

func StrToInt(s string) int {
	
	if len(s) == 0 { // check first
		return 0
	}
	
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

	if start >= len(s) { // handle string like "+" or "-" only
		return 0
	}

	for i := start; i < len(s); i++ {	//loop acording string length
		if s[i] < '0' || s[i] > '9' { // check If the string is not a valid number
			return 0
		}	
		digit := s[i] - '0'        // convert character to integer
		num = num*10 + int(digit)  // shift previous digits left and add new
	}

	return sign * num
}

func BulkAtoi(arr []string) []int {
	var result []int 
	for i:=0; i<len(arr); i++ {
		result = append(result, StrToInt(arr[i]))
	}
	return result
}

/* func main() {
	fmt.Println(BulkAtoi([]string{"8", "kood", "-13"}))
} */