package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("get:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	twoD = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2d: ", twoD)
}

// here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array's type. By default an array is zero-valued, which for ints means 0s.
// we can set a value at an index using the array[index] = value syntax, and get a value with array[index].
// the builtin len returns the length of an array.
// use this syntax to declare and initialize an array in one line.
// use this syntax to declare and initialize an array in one line.
// array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
// note that arrays appear in the format [v1 v2 v3 ...] when printed with fmt.Println.
