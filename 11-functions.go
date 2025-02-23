package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}

// here's a function that takes two ints and returns their sum as an int.
// Go requires explicit returns, i.e. it won't automatically return thee value of the last expression
//when you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed-
//+ parameters up to the final parameeter that declares the type.
//call a function just as you'd expect, with name(args).
