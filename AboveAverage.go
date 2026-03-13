package main

import "fmt"

type Item struct {
	Name  string
	Value float64
}

func AboveAverage(items []Item) []string {
	var total float64
	for _, item := range items {
		total += item.Value
	}
	average := total / float64(len(items))

	var result []string
	for _, item := range items {
		if item.Value > average {
			result = append(result, item.Name)
		}
	}
	return result
}

func main() {
	items := []Item{
    {"A", 10},
    {"B", 30},
    {"C", 20},
    {"D", 50},
}
	fmt.Println("Items above average:", AboveAverage(items))
}