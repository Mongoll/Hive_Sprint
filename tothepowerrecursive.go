package main
/* 
import "fmt" */

func ToThePowerRecursive(n int, power int) int {
	if power == 0 {
		return 1
	} else if power > 0 {
		return n * ToThePowerRecursive(n, power-1)
	}
	return 0

}
/* func main() {
	fmt.Println(ToThePowerRecursive(2, 10))
} */