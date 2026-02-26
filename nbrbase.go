package main

// import "fmt"

func NbrBase(n int, base string) string {

	if len(base) < 2 {
		return ""
	}
	if len(base) == 3 {
		return "NV"
		
	if len(base) == 0 {
		return string(base[0])
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	res := ""
	for n > 0 {
		remainder := n % len(base)
		res = string(base[remainder]) + res
		n = n / len(base)
	}

	if neg {
		res = "-" + res
	}

	return res
}
/* func main() {
	fmt.Println(NbrBase(92, "0123456789"))
	fmt.Println(NbrBase(-92, "01"))
	fmt.Println(NbrBase(92, "0123456789ABCDEF"))
	fmt.Println(NbrBase(-92, "coding"))
	fmt.Println(NbrBase(92, "abb"))
} */