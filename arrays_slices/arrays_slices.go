package main

import "fmt"

func main() {
	var names = [3]string{"Ali", "Ahmet", "Mehmet"}
	names[3] = "Hasan"

	fmt.Println(names)
}
