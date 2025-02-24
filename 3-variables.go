package main

import "fmt"

func main() {
	var a = "initial" // var ile 1 veya daha fazla değişken tanımlayabiliriz
	fmt.Println(a)
	var b, c int = 1, 2 // birden fazla değişken de tanımlayabiliriz.
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int // karşılık belirtilmediğinde değeri sıfırdır.
	fmt.Println(e)

	f := "apple" // := sözdizimi, değişken tanımlamak için kullanılır.
	fmt.Println(f)
}
