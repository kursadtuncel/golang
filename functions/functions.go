package main

import "fmt"

func main() {
	total := add(10, 20)
	fmt.Println(total)

	//total, difference, multiply := calculation(10,20)
}

/*
func calculation(x int, y int) (int, int, int) {
return x + y, x - y, x * y
}
*/

func add(x int, y int) int {
	//fmt.Println(x + y)
	return x + y
}
