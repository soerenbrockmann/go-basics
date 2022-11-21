// package main

// import "fmt"

// type Welcome struct {
// 	Greet string
// }

// type InterfaceForMethods interface {
// 	Method()
// }

// // This method means type T implements the interface I,
// // but we don't need to explicitly declare that it does so.
// func (t Welcome) Method() {
// 	fmt.Println(t.Greet)
// }

// func main() {
// 	var i  InterfaceForMethods = Welcome{"hello"}
// 	i.Method()
// }



// Interface values
// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)
// An interface value holds a value of a specific underlying concrete type.
// Calling a method on an interface value executes the method of the same name on its underlying type.


// package main

// import (
// 	"fmt"
// 	"math"
// )

// type InterfaceForMethods interface {
// 	Method()
// }

// type Welcome struct {
// 	Greet string
// }

// func (t *Welcome) Method() {
// 	fmt.Println(t.Greet)
// }

// type F float64

// func (f F) Method() {
// 	fmt.Println(f)
// }

// func describe(i InterfaceForMethods) {
// 	fmt.Printf("(%v, %Welcome)\n", i, i)
// }

// func main() {
// 	var i InterfaceForMethods

// 	i = &Welcome{"Hello"}
// 	describe(i)
// 	i.Method()

// 	i = F(math.Pi)
// 	describe(i)
// 	i.Method()
// }



// Interface values with nil underlying values
// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

// In some languages this would trigger a null pointer exception, but in Go it is common to write methods 
// that gracefully handle being called with a nil receiver (as with the method M in this example.)

// Note that an interface value that holds a nil concrete value is itself non-nil.


// package main

// import "fmt"

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// func (t *T) M() {
// 	if t == nil {
// 		fmt.Println("<nil>")
// 		return
// 	}
// 	fmt.Println(t.S)
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// func main() {
// 	var i I

// 	fmt.Println(i)

// 	var t *T

// 	fmt.Println(t)

// 	i = t

// 	fmt.Println(i)
// 	describe(i)
// 	i.M()

// 	i = &T{"hello"}
//  	describe(i)
//  	i.M()
//  }

// Nil interface values
// A nil interface value holds neither value nor concrete type.

// Calling a method on a nil interface is a run-time error because 
// there is no type inside the interface tuple to indicate which concrete method to call.

// package main

// import "fmt"

// type I interface {
// 	M()
// }


// func main() {
// 	var i I
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }


// The empty interface
// The interface type that specifies zero methods is known as the empty interface:

// interface{}
// An empty interface may hold values of any type. (Every type implements at least zero methods.)

// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print 
// takes any number of arguments of type interface{}.

// package main

// import "fmt"

// func main() {
// 	var i interface{}
// 	describe(i)

// 	i = 24.20
// 	describe(i)

// 	i = "hello"
// 	describe(i)
// }

// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// Type assertions
// A type assertion provides access to an interface value's underlying concrete value.

// t := i.(T)
// This statement asserts that the interface value i holds the concrete type T and assigns 
//the underlying T value to the variable t.

// If i does not hold a T, the statement will trigger a panic.

// To test whether an interface value holds a specific type, a type assertion can return two values: 
// the underlying value and a boolean value that reports whether the assertion succeeded.

// t, ok := i.(T)
// If i holds a T, then t will be the underlying value and ok will be true.

// If not, ok will be false and t will be the zero value of type T, and no panic occurs.

// Note the similarity between this syntax and that of reading from a map.

// < 15/26 >
// type-assertions.go Syntax Imports
// package main

// import "fmt"

// func main() {
// 	var i interface{} = "hello"

// 	s := i.(string)
// 	fmt.Println(s)

// 	s, ok := i.(string)
// 	fmt.Println(s, ok)

// 	f, ok := i.(float64)
// 	fmt.Println(f, ok)

// 	f = i.(float64) // panic
// 	fmt.Println(f)
// }

// Type switches
// A type switch is a construct that permits several type assertions in series.

// A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

// switch v := i.(type) {
// case T:
//     // here v has type T
// case S:
//     // here v has type S
// default:
//     // no match; here v has the same type as i
// }
// The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.

// This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.

// < 16/26 >
// type-switches.go Syntax Imports

// package main

// import "fmt"

// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	case string:
// 		fmt.Printf("%q is %v bytes long\n", v, len(v))
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)
// 	}
// }

// func main() {
// 	do(21)
// 	do("hello")
// 	do(true)
// }


// Stringers
// One of the most ubiquitous interfaces is Stringer defined by the fmt package.

// type Stringer interface {
//     String() string
// }
// A Stringer is a type that can describe itself as a string. 
// The fmt package (and many others) look for this interface to print values.

// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v yearssss)", p.Name, p.Age)
// }

// func main() {
// 	a := Person{"Arthur Dent", 42}
// 	z := Person{"Zaphod Beeblebrox", 9001}
// 	fmt.Println(a, z)
// }


// Errors
// Go programs express error state with error values.

// The error type is a built-in interface similar to fmt.Stringer:

// type error interface {
//     Error() string
// }

// (As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

// Functions often return an error value, and calling code should handle errors by testing 
// whether the error equals nil.

// i, err := strconv.Atoi("42")
// if err != nil {
//     fmt.Printf("couldn't convert number: %v\n", err)
//     return
// }
// fmt.Println("Converted integer:", i)
// A nil error denotes success; a non-nil error denotes failure.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type MyError struct {
// 	When time.Time
// 	What string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s",
// 		e.When, e.What)
// }

// func run() error {
// 	return &MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// }

// func main() {
// 	// Long version
// 	// err := run()
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// Short version
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }

// Readers
// The io package specifies the io.Reader interface, which represents the read end of a stream of data.

// The Go standard library contains many implementations of this interface, including files, network connections, 
// compressors, ciphers, and others.

// The io.Reader interface has a Read method:

// func (T) Read(b []byte) (n int, err error)
// Read populates the given byte slice with data and returns the number of bytes populated and an error value.
// It returns an io.EOF error when the stream ends.

// The example code creates a strings.Reader and consumes its output 8 bytes at a time.

// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	r := strings.NewReader("Hello, Reader!")

// 	b := make([]byte, 8)
// 	for {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }


// Images
// Package image defines the Image interface:

// package image

// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }

// Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, 
// as the declaration is inside package image.

// (See the documentation for all the details.)

// The color.Color and color.Model types are also interfaces, but we'll ignore that by 
// using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and 
// types are specified by the image/color package

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

