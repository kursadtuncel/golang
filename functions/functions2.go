package main

import "fmt"

func main() {
	var numbers = []int{10, 20, 2, 3, 5}

	sum(numbers)

	fmt.Println(sum(numbers))
}

func sum(numbers []int) int {
	//foreach ile
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
