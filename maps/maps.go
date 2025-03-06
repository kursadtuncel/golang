package main

import "fmt"

func main() {
	/*
		var names map[string]int

		names = make(map[string]int, 0)

		names["Ali"] = 1
		names["Ahmet"] = 2
		names["Mehmet"] = 3

		fmt.Println(names["Ali"])
	*/

	names := map[string]int{
		"Ali":    1,
		"Ahmet":  2,
		"Mehmet": 3,
	}

	// delete(names,"Ahmet")

	fmt.Println(names)
}
