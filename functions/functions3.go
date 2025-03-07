package main

import "fmt"

func main() {
	fmt.Println(sum(3, 4, 5))
	fmt.Println(sum(3, 4, 5, 6))

}

func sum(numbers ...int) int { // "..." sayesinde istediğimiz kadar parametre ekleyebiliriz
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
