package main

/* import "fmt" */

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointDiff(p1, p2 Point) Point {
	if p1.X > p2.X {
		return p1
	} else if p1.X == p2.X {
		if p1.Y > p2.Y {
			return p1
		}
	}
	return p2
}

/* func main() {
	fmt.Println(PointDiff(Point{3, 4, "text"}, Point{3, 5, "text"}))
} */