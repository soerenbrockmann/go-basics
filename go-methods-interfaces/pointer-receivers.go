package main

import (
	"fmt"
	"math"
)

type Tensor struct {
	X, Y float64
}

// As normal function
// func MaxValue(t Tensor) float64 {
// 	return math.Max(t.X, t.Y)
// }

func (t Tensor) MaxValue() float64 {
	return math.Max(t.X, t.Y)
}

// Methods with pointer receivers can modify the value to which the receiver points 
// (as Scale does here). Since methods often need to modify their receiver, pointer 
// receivers are more common than value receivers.

// As normal function
// func Scale(t *Tensor, s float64) {
// 	t.X = t.X * s
// 	t.Y = t.Y * s
// }

func (t *Tensor) Scale(s float64) {
	t.X = t.X * s
	t.Y = t.Y * s
}

// Comparing the previous two programs, you might notice that functions with a pointer 
// argument must take a pointer. while methods with pointer receivers take either a value 
// or a pointer as the receiver when they are called:
func main(){
	t := Tensor{3, 10}
	t.Scale(10)

	maxValue := t.MaxValue()
	fmt.Println(maxValue)
}

// Choosing a value or pointer receiver
// There are two reasons to use a pointer receiver.

// The first is so that the method can modify the value that its receiver points to.

// The second is to avoid copying the value on each method call. 
// This can be more efficient if the receiver is a large struct, for example.