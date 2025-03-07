package main

import "fmt"

func main() {
	var condition = true

	defer cleanup()
	if condition {
		panic("an error occured")
	}
}

func cleanup() {
	fmt.Println("cleanup worked...")
}
