package main

import "fmt"

var a int
var b = "HOt"

func main() {
	a = int(b)
	fmt.Println("Hi", a)
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println((i))
		}
	}
}

func foo() {
	fmt.Println("Fu***")
}
