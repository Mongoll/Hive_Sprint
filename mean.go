package sprint

import "fmt"

func Mean(a, b, c float32) float32 {
	return (a+b+c)/3.0 //can use just 3, but it's float and better use with '.'
}

func main() {
	fmt.Println(Mean(5.0, 10.0, 3.5))
}