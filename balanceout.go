package main

import "fmt"

func BalanceOut(arr []bool) []bool {
	trueCount := 0
	falseCount := 0
	for _, value := range arr {
		if value {
			trueCount++
		} else {
			falseCount++
		}
	}
	for trueCount < falseCount {
		arr = append(arr, true)
		trueCount++
	} 
	for trueCount > falseCount {
		arr = append(arr, false)
		falseCount++
	}
	return arr
}

func main(){
	fmt.Println(BalanceOut([]bool{true, false, false, false}))
}