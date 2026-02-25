
import "fmt"

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointText(p Point) Point {
	p.Text = "Text in (" + fmt.Sprintf("%v", p.X) + ", " + fmt.Sprintf("%v", p.Y) + ")"
	return p
}

func main() {
	fmt.Println(PointText(Point{3, 5, "text"}))
}
