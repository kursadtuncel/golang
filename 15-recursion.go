package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

}

//this fact function calls itself until it reaches the base case of fact(0).
//anonymous functions can also be recursive, but this requires explicity declaring a variable with-
//+ var to store the function before it's defined.

//since fib was previously declared in main, go knows which function to call with fib here.
