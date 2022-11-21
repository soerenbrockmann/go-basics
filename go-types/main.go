package main

import (
    "fmt"
)

func runArrays() {
    var names [2]string
    names[0] = "Hannes"
    names[1] = "Peter"
    fmt.Println(names)
    fmt.Println(names[0])
    fmt.Println(names[1])

    numbers := [6]int{1, 2, 3, 4, 5, 6}
    fmt.Println(numbers)
}

// Slices are like references to arrays
// A slice does not store any data, it just describes
// a section of an underlying array.
// Changing the elements of a slice modifies the corresponding
// elements of its underlying array.

// Other slices that share the same underlying array will see those
// changes.

func runSlices() {
    // Basic Slices
    names := [4]string{
        "Hannes",
        "Peter",
        "Gary",
        "Harry",
    }
    fmt.Println(names)
    var snames []string = names[1:4]
    fmt.Println(snames)

    // Changing the elements of a slice modifies the corresponding
    // elements of its underlying array.

    a := names[0:2]
    b := names[2:4]
    fmt.Println(a, b)

    b[0] = "SECRET"
    fmt.Println(a, b)
    fmt.Println(names)

    // Slice literals
    // A slice literal is like an array literal without the length.
    numbers := []int{1, 2, 3, 4, 5, 6}
    fmt.Println(numbers)

    boolValues := []bool{true, false, true, true, false, true}
    fmt.Println(boolValues)

    structValues := []struct {
        index int
        value bool
    }{
        {0, true},
        {1, false},
        {2, true},
        {3, true},
        {4, false},
        {5, true},
    }
    fmt.Println(structValues)

    // Slice defaults
    // When slicing, you may omit
    // the high or low bounds to use their defaults instead.
    // defaults := []int{2, 3, 5, 7, 11, 13}

    // defaults = defaults[1:4]
    // fmt.Println(defaults)

    // defaults = defaults[:2]
    // fmt.Println(defaults)

    // defaults = defaults[1:]
    // fmt.Println(defaults)

    // Slice length and capacity
    // The length of a slice is the number of elements it contains.
    // The capacity of a slice is the number of elements in the underlying
    // array, counting from the first element in the slice.

    // The length and capacity of a slice s can be obtained using the
    // expressions len(s) and cap(s).
    // s := []int{2, 3, 5, 7, 11, 13}
    // fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // // Slice the slice to give it zero length.
    // s = s[:0]
    // fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // // Extend its length.
    // s = s[:4]
    // fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // // Drop its first two values.
    // s = s[2:]
    // fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // Nil slices
    // The zero value of a slice is nil.

    // A nil slice has a length and capacity of 0 and has no
    // underlying array.
    // var snil []int
    // fmt.Println(snil, len(snil), cap(snil))
    // if snil == nil {
    //  fmt.Println("nil!")
    // }

    // Creating a slice with make
    // Slices can be created with the built-in make function; this is
    // how you create dynamically-sized arrays.
    // amake := make([]int, 5)
    // fmt.Printf("%s len=%d cap=%d %v\n",
    //  "amake", len(amake), cap(amake), amake)

    // bmake := make([]int, 0, 5)
    // fmt.Printf("%s len=%d cap=%d %v\n",
    //  "bmake", len(bmake), cap(bmake), bmake)

    // cmake := bmake[:2]
    // fmt.Printf("%s len=%d cap=%d %v\n",
    //  "cmake", len(cmake), cap(cmake), cmake)

    // dmake := cmake[2:5]
    // fmt.Printf("%s len=%d cap=%d %v\n",
    //  "dmake", len(dmake), cap(dmake), dmake)

    // Slices of slices
    // Slices can contain any type, including other slices.
    // Create a tic-tac-toe board.
    // board := [][]string{
    //  []string{"_", "_", "_"},
    //  []string{"_", "_", "_"},
    //  []string{"_", "_", "_"},
    // }

    // // The players take turns.
    // board[0][0] = "X"
    // board[2][2] = "O"
    // board[1][2] = "X"
    // board[1][0] = "O"
    // board[0][2] = "X"

    // for i := 0; i < len(board); i++ {
    //  fmt.Printf("%s\n", strings.Join(board[i], " "))
    // }

    // Appending to a slice
    // It is common to append new elements to a slice, and so
    // Go provides a built-in append function. The documentation
    // of the built-in package describes append.
    var s []int
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // append works on nil slices.
    s = append(s, 0)
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // The slice grows as needed.
    s = append(s, 1)
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // We can add more than one element at a time.
    s = append(s, 2, 3, 4)
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
    //  // Range
    //  // The range form of the for loop iterates over a slice or map.
    //  var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

    //  for i, v := range pow {
    //      fmt.Printf("2**%d = %d\n", i, v)
    //  }

    //  //  Range continued
    //  // You can skip the index or value by assigning to _.
    //  pow := make([]int, 10)
    //  for i := range pow {
    //      pow[i] = 1 << uint(i) // == 2**i
    //  }
    //  for _, value := range pow {
    //      fmt.Printf("%d\n", value)
    //  }
    // }

}

func runMaps() {

    // Map literals
    // Map literals are like struct literals, but the keys are required.
    type Person struct {
        name string
        age  int
    }

    record := map[string]Person{
        "Company X": Person{
            "Sören", 25,
        },
        "Company Y": Person{
            "Hannes", 30,
        },
    }

    fmt.Println("New record", record)

    // Map literals continued
    // If the top-level type is just a type name, you can omit it
    // from the elements of the literal.

    record = map[string]Person{
        "Company X": {"Sören", 25},
        "Company Y": {"Hannes", 30},
    }
    fmt.Println("New record short", record)

    // Mutating Maps
    person := make(map[string]int)
    person["age"] = 20
    fmt.Println("Age is:", person["age"])

    person["age"] = 30
    fmt.Println("New age is", person["age"])

    delete(person, "age")
    fmt.Println("Age deleted:", person["age"])

    // v, ok := person["age"]
    // fmt.Println("The value:", v, "Present?", ok)

}

func main() {
    // runArrays()
    // runSlices()
    runMaps()
}

