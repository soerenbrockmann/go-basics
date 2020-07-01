package main

import "fmt"

func getFullName(firstName string, lastName string) func() string {

	return func() string {
		return firstName + " " + lastName
	}

}

func main() {
	myFunc := getFullName("Sören", "Brockmann")

	fmt.Println(myFunc())
}
