package main

/* import "fmt" */

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	pi := float32(3.14)
	dia := r * 2
	area := pi * (r * r)
	perimeter := 2 * pi * r
	return Circle{Radius: r, Diameter: dia, Area: area, Perimeter: perimeter}
}
/* func main() {
	fmt.Println(GetCircle(5))
}
 */