package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	youtube := Person{name: "Sören", age: 24}
	fmt.Println(youtube.name)
}
