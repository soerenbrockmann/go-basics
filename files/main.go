package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	myData := []byte("Great job!")
	err := ioutil.WriteFile("data.data", myData, 0777)
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("data.data")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	f, err := os.OpenFile("data.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("data.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}
