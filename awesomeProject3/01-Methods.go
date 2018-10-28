package awesomeProject3

import (
	"fmt"
	"math"
)

// struct
type Vertex struct {
	X, Y float64
}

// special argument to define function on vertex
// this is called method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())
}
