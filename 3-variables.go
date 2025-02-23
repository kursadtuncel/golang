package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

// var declares 1 or more variables
// you can declares 1 or more variables
// you can declare multiple variables at once
//go will infer the type of initialized variables
// variablees declared without a corresponding initialization are zero-valued.
// + for example, the zero value for an int is 0
// the := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case
// this syntax is only available inside functions
