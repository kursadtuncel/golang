package main

import "fmt"

func main() {
	/*
		var a int

		var p *int //pointer

		a = 10

		p = &a // a'nın adres değerini p'ye depolamış olduk.
		fmt.Println(a)

		*p = 20 // pointerin gösterdiği adresteki değeri 20 olarak değiştirdik.

		fmt.Println(a)
	*/

	var a = 10
	var b int
	var p *int

	b = a
	p = &a

	*p = 20
	fmt.Println(a, b)
}
