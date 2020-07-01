package main

import "fmt"

var name *string

func init() {
	// name = "Sören"
	fmt.Println("I am getting called immediately. Setting name to", &name)

}
func main() {
	fmt.Println("Main print")
}
