package main

import (
	"fmt"
	"math"
)

type Tensor struct {
	X, Y float64
}

func (t Tensor) MaxValue() float64 {
	return math.Max(t.X, t.Y)
}

// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// You can only declare a method with a receiver whose type is defined in the same package as the method. 
// You cannot declare a method with a receiver whose type is defined in another package (which includes 
// the built-in types such as int).
func main() {
	t := Tensor{3, 10}

	maxValue := t.MaxValue()

	fmt.Println(maxValue)
}