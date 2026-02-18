package main

import "fmt"

// IsNegative returns true if n is less than 0, otherwise false.
func IsNegative(n int) bool {
    return n < 0
}

func main() {
    // Examples
    fmt.Println(IsNegative(6))  // false
    fmt.Println(IsNegative(-5)) // true
    fmt.Println(IsNegative(0))  // false
}