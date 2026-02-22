package main

import "fmt"
import "math"

func Casting(n float64) int {
	return int(math.Round(n)) //round float to nearest int
}

func main() {
	fmt.Println(Casting(4.25))
}