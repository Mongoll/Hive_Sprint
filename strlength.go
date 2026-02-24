package main

func StrLength(s string) []int {
	var arr []int
	arr = append(arr, len([]rune(s)), len(s))
	return arr
}