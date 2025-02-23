package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { //this area method has a receiver type of *rect
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("are: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

//methods can be defined for either pointer or value reeceiver types.
//here's an example of a value receiver.

//go automatically handles conversion beetween values and pointers for method calls.
//you may want to use a pointer receiver type of avoid copying on method calls or to allow the-
//+ method to mutate the receeiving struct.
