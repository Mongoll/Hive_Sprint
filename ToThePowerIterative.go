package main

import "fmt"

func ToThePowerIterative(n int, power int) int {
	
	f := 1
	if power > 0 {
		for i := 1; i <= power; i++ {
			f *= n
		}
	} else {
		f = 0
	}
	return f

}
func main() {
	fmt.Println(ToThePowerIterative(2, 10))
}
