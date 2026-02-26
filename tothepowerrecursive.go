package main

import "fmt"

func ToThePowerIterative(n int, power int) int {
	if power == 0 {
		return 1
	} else if power > 0 {
		return n * ToThePowerIterative(n, power-1)
	}
	return 0

}
func main() {
	fmt.Println(ToThePowerIterative(2, 10))
}