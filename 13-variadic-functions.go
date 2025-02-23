package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

//here's a function that will take an arbitrary number of ints as arguments.
//within the function, the type of nums is equivalent to [] int. we can call len(nums), iterate over it with range, etc.
//variadic functions can be called in the usual way with individual arguments.
//if you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
