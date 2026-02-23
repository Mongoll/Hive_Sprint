package main

// import "fmt"

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
if from > to {
	from, to = to, from
}
if from < 0 {
	from = 0
}
if to > len(arr) {
	to = len(arr)
}
if from >= len(arr) || from >= to {
		return arr
	}
	return append(arr[:from], arr[to:]...)
}

/* func main() {
	fmt.Println(RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1))
} */