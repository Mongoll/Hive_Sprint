package sprint

/* import (
    "fmt"
) */

// BetweenLimits returns a string containing all runes between 'from' and 'to' (exclusive)
func BetweenLimits(from, to rune) string {
    // Make sure from < to
    if from > to {
        from, to = to, from
    }

    result := ""
    for r := from + 1; r < to; r++ {
        result += string(r)
    }
    return result
}
/* 
func main() {
    fmt.Println(BetweenLimits('j', 'f')) // "ghi"
    fmt.Println(BetweenLimits('f', 'j')) // "ghi"
    fmt.Println(BetweenLimits('a', 'd')) // "bc"
    fmt.Println(BetweenLimits('z', 'x')) // "y"
} */