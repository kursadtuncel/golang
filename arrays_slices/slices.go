package main

import "fmt"

func main() {
	var names = []string{"Ali", "Ahmet", "Mehmet"}

	names = append(names, "Ferhat")
	fmt.Println(names)
}
