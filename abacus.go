package main
import [fmt]
func Abacus(a int, b int) int {
	if  b == 0 {
		panic("division by zero")
	}
	return a / b
}
func main() {
	fmt.Println(Abacus(10,2))
}