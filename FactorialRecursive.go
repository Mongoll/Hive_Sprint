package main

/* import "fmt" */

func FactorialIterative(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * FactorialIterative(n-1)
}

/* func main() {
	fmt.Println(FactorialIterative(5))
} */