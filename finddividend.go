package sprint

import (
    "fmt"
)

// FindDividend returns the first number in [from, to) divisible by divisor.
// Returns -1 if no such number exists or if divisor is 0.
func FindDividend(from, to, divisor int) int {
    if divisor == 0 {
        return -1 // avoid division by zero
    }

    for i := from; i < to; i++ {
        if i%divisor == 0 {
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(FindDividend(5, 17, 4))  // 8
    fmt.Println(FindDividend(10, 18, 9)) // -1
    fmt.Println(FindDividend(1, 10, 3))  // 3
    fmt.Println(FindDividend(5, 10, 0))  // -1 (invalid divisor)
}