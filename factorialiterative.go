package main

import "fmt"

func FactorialIterative(n int) int {
	var f int
	f = 1
	if n>0{
		for i := 1; i <= n; i++ {
			f *= i
		}
	}
	return f
}

func main() {
	fmt.Println(FactorialIterative(5))
}