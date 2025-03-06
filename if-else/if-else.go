package main

import "fmt"

func main() {
	/*var age = 19

	if age >= 18 {
		fmt.Println("you can vote")
	} else {
		fmt.Println("you can't vote")
	} */

	a := 10
	b := 20
	c := 30

	if a >= b && a >= c {
		fmt.Println("en buyuk deger a'dır")
	} else if b >= a && b >= c {
		fmt.Println("en buyuk deger b'dir")
	} else {
		fmt.Println("en buyuk deger c'dir")
	}
}
