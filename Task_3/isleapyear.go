package main

import (
    "fmt"
)
 
// IsLeapYear returns true if the given year is a leap year, false otherwise
func IsLeapYear(year int) bool {
    if year%4 != 0 {
        return false
    } else if year%100 != 0 {
        return true
    } else if year%400 != 0 {
        return false
    } else {
        return true
    }
}

func main() {
    fmt.Println(IsLeapYear(2020)) // true
    fmt.Println(IsLeapYear(1900)) // false
    fmt.Println(IsLeapYear(2000)) // true
    fmt.Println(IsLeapYear(2021)) // false
} 