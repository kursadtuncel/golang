package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// the most basic type, with a single condition
// a classic initial/condition/after for loop
// another way of accompishing the basic "do this N times"-
// + iteration is range over an integer
// for without a condition will lopp repeatedly until you break out of the-
// + loop or return from the enclosing function
// you can also continue to the next iteration of the loop
