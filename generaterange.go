package main

// import "fmt"

func GenerateRange(min, max int) []int {
	nums := make([]int, 0)
	if max > min {
		for i := min; i < max; i++ {
			nums = append(nums, i)
		}
	}
	return nums
}

/* func main() {
	fmt.Println(GenerateRange(-1, 4))

} */