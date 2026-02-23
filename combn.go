package main

import "fmt"

func CombN(n int) []string {
	var comb []string
	var recurs func(from int, current string)
	recurs = func(from int, current string) {
		if len(current) == n {
			comb = append(comb, current)
			return
		}
		for i := from; i <= 9; i++ {
			recurs(i+1, current+string(rune('0'+i)))
		}
	}
	recurs(0, "")
	return comb
}

func main() {
	fmt.Println(CombN(9))
}