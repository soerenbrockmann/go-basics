package main

import (
	"container/list"
	"fmt"
)

func main() {

	mylist := list.New()
	mylist.PushBack(1)
	mylist.PushFront(2)

	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		if element.Value != 1 {
			mylist.Remove(element)
		}
	}

	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
	}

}
