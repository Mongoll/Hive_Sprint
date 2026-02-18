package main

import "fmt"


// Accumulate returns the sum of all integers from 0 to n (inclusive) if n is positive.
// Returns 0 if n is negative.
func Accumulate(n int) int {
    if n < 0 {
        return 0
    }

    sum := 0
    for i := 0; i <= n; i++ {
        sum += i
    }
    return sum
}

func main() {
    fmt.Println(Accumulate(4))   // 10
    fmt.Println(Accumulate(0))   // 0
    fmt.Println(Accumulate(-5))  // 0
    fmt.Println(Accumulate(10))  // 55
} 