package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

// the (int, int) in this function signature shows that the function returns 2 ints
// here we use the 2 different return values from the call with multiple assignment.

//if you only want a subset of the returned values, use the blank identifier _.
