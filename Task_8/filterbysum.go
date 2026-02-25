package main

// import "fmt"

func FilterBySum(arr [][]int, limit int) [][]int {
	var result [][]int
	for _, subArray := range arr {
		sum := 0
		for _, num := range subArray {
			sum += num
		}
		if sum >= limit {
			result = append(result, subArray)
		}
	}
	return result
}

/* func main() {
	fmt.Println(FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 9))
} */