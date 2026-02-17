package sprint

/* import (
    "fmt"
) */

func IntVsFloat(i int, f float32) string {
    fi := float32(i)
    
    if fi > f {
        return "Integer"
    } else if fi < f {
        return "Float"
    } else {
        return "Same"
    }
}

/* func main() {
    fmt.Println(IntVsFloat(5, 8.8))   // "Float"
    fmt.Println(IntVsFloat(10, 3.5))  // "Integer"
    fmt.Println(IntVsFloat(7, 7.0))   // "Same"
} */