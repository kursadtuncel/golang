package main

import "fmt"

var x = 10

func main() {
	fmt.Println(x)
	test()
	fmt.Println(x)

}

func test() {
	var x = 20 //function scope
	fmt.Println(x)
}
