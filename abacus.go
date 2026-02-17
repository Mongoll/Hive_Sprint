package main

func Abacus(a int, b int) int {
	if  b == 0 {
		panic("division by zero")
	}
	return a / b
}
