package main

import "fmt"

type Item struct {
	Name  string
	Value float64
}

func main() {
	items := []Item{
		{"A", 0.05},
		{"B", 0.1},
		{"C", 0.8},
	}

	maxItem := items[0]

	for _, item := range items {
		if item.Value > maxItem.Value {
			maxItem = item
		}
	}

	fmt.Println("Name with max value:", maxItem.Name)
}