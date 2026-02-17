package sprint

/* import (
    "fmt"
) */

// Doop performs basic arithmetic operations and handles invalid input
func Doop(a int, op string, b int) int {
    switch op {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        if b == 0 {
            return 0 // avoid division by zero
        }
        return a / b
    case "%":
        if b == 0 {
            return 0 // avoid modulo by zero
        }
        return a % b
    default:
        return 0 // invalid operator
    }
}
/* 
func main() {
    fmt.Println(Doop(5, "+", 3))   // 8
    fmt.Println(Doop(8, "/", 0))   // 0
    fmt.Println(Doop(10, "%", 3))  // 1
    fmt.Println(Doop(7, "-", 2))   // 5
    fmt.Println(Doop(4, "x", 2))   // 0 (invalid operator)
} */