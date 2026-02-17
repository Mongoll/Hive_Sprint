package sprint

/* import (
    "fmt"
) */

// CountDivisible counts the numbers in the range [from, to) with a given step
// that are divisible by divisor. Returns 0 if step <= 0 or divisor == 0.
func CountDivisible(from, to, step, divisor int) int {
    if step <= 0 || divisor == 0 {
        return 0
    }

    count := 0
    for i := from; i < to; i += step {
        if i%divisor == 0 {
            count++
        }
    }
    return count
}
/* 
func main() {
    fmt.Println(CountDivisible(5, 17, 2, 3))  // 2
    fmt.Println(CountDivisible(1, 10, 1, 2))  // 5
    fmt.Println(CountDivisible(5, 17, 0, 3))  // 0 (step = 0)
    fmt.Println(CountDivisible(5, 17, 2, 0))  // 0 (divisor = 0)
} */