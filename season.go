package main

/* import (
    "fmt"
    "strings"
) */

// Season returns the season name for a given month abbreviation (lowercase 3 letters)
func Season(month string) string {
    // Normalize input to lowercase
    m := strings.ToLower(month)

    switch m {
    case "jan", "feb", "dec":
        return "winter"
    case "mar", "apr", "may":
        return "spring"
    case "jun", "jul", "aug":
        return "summer"
    case "sep", "oct", "nov":
        return "autumn"
    default:
        return "invalid input: " + month
    }
}
/* 
func main() {
    fmt.Println(Season("feb"))        // winter
    fmt.Println(Season("September"))  // invalid input: September
    fmt.Println(Season("Jul"))        // summer
    fmt.Println(Season("abc"))        // invalid input: abc
} */