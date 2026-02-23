package main

// import "fmt"

func SortIntegerTable(table []int) []int {
	n := len(table)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	return table
}

/* func main() {
	fmt.Println(SortIntegerTable([]int{2, 0, 5, 4, 1, 3}))
} */