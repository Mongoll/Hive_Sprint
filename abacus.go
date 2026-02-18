package main

import "fmt"

func Abacus(a int, b int) int {
	if  b == 0 {
		panic("division by zero")
	}
	return a / b
}

func main() {
	result := Abacus(10, 2)
	fmt.Println("Result:", result)
} 